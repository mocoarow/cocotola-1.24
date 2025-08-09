import * as client from "openid-client";
import type { LoaderFunctionArgs, MetaFunction } from "react-router";
import { Link, useLoaderData } from "react-router";
import { getAuthConfig } from "~/auth.server";
import { Button } from "~/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";
import { commitLoginDataSession, getLoginDataSession } from "~/session.server";
export const meta: MetaFunction = () => {
  return [{ title: "Login" }];
};

export async function loader({ request }: LoaderFunctionArgs) {
  console.log("_auth.login.tsx::loader");

  const config: client.Configuration = await getAuthConfig();
  const codeVerifier = client.randomPKCECodeVerifier();
  const nonce = client.randomNonce();
  const state = client.randomState();

  const session = await getLoginDataSession(request);

  const parameters: Record<string, string> = {
    redirect_uri: "http://localhost:5173/callback",
    scope: "openid profile email",
    code_challenge: await client.calculatePKCECodeChallenge(codeVerifier),
    code_challenge_method: "S256",
    nonce: nonce,
    state: state,
  };

  session.set("codeVerifier", codeVerifier);
  session.set("nonce", nonce);
  session.set("state", state);

  const redirectTo = client.buildAuthorizationUrl(config, parameters);
  const googleAuthUrl = redirectTo.toString();

  return new Response(JSON.stringify({ googleAuthUrl: googleAuthUrl }), {
    headers: {
      "Content-Type": "application/json",
      "Set-Cookie": await commitLoginDataSession(session),
    },
  });
}

export default function Index() {
  const { googleAuthUrl } = useLoaderData<typeof loader>();
  // console.log(googleAuthUrl);
  return (
    <div className="flex min-h-svh flex-col items-center justify-center gap-6 bg-muted p-6 md:p-10">
      <div className="flex w-full max-w-sm flex-col gap-6">
        <Card>
          <CardHeader className="text-center">
            <CardTitle className="text-xl">Welcome back</CardTitle>
            <CardDescription>Login with your Google account</CardDescription>
          </CardHeader>
          <CardContent>
            <Link to={googleAuthUrl}>
              <Button variant="outline" className="w-full">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                  <title>Login</title>
                  <path
                    d="M12.48 10.92v3.28h7.84c-.24 1.84-.853 3.187-1.787 4.133-1.147 1.147-2.933 2.4-6.053 2.4-4.827 0-8.6-3.893-8.6-8.72s3.773-8.72 8.6-8.72c2.6 0 4.507 1.027 5.907 2.347l2.307-2.307C18.747 1.44 16.133 0 12.48 0 5.867 0 .307 5.387.307 12s5.56 12 12.173 12c3.573 0 6.267-1.173 8.373-3.36 2.16-2.16 2.84-5.213 2.84-7.667 0-.76-.053-1.467-.173-2.053H12.48z"
                    fill="currentColor"
                  />
                </svg>

                <span>Login with Google</span>
              </Button>
            </Link>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
