import * as client from "openid-client";
import type { MetaFunction } from "react-router";
import { type LoaderFunctionArgs, redirect } from "react-router";

import { getAuthConfig } from "~/auth.server";
import {
  commitUserDataSession,
  destroyLoginDataSession,
  getLoginDataSession,
  getUserDataSession,
} from "~/session.server";

export const meta: MetaFunction = () => {
  return [{ title: "Callback" }];
};

export async function loader({ request }: LoaderFunctionArgs) {
  console.log("_auth.callback.tsx::loader");

  const config: client.Configuration = await getAuthConfig();
  const loginDataSession = await getLoginDataSession(request);
  const sessionUserData = await getUserDataSession(request);

  const codeVerifier = loginDataSession.get("codeVerifier");
  const nonce = loginDataSession.get("nonce");
  const state = loginDataSession.get("state");

  const currentUrl = new URL(request.url);

  try {
    const tokens = await client.authorizationCodeGrant(config, currentUrl, {
      pkceCodeVerifier: codeVerifier,
      expectedNonce: nonce,
      expectedState: state,
      idTokenExpected: true,
    });

    const { access_token: accessToken } = tokens;
    const claims = tokens.claims();
    if (!claims) {
      console.log("Token Endpoint Response", tokens);
      throw new Error("failed to get claims");
    }
    const { sub } = claims;

    sessionUserData.set("accessToken", accessToken);
    sessionUserData.set("sub", sub);

    if (!accessToken || !sub) {
      console.log("ID Token Claims", claims);
      throw new Error("accessToken and sub are required");
    }

    const userInfo = await client.fetchUserInfo(config, accessToken, sub);
    if (!userInfo.email || !userInfo.name) {
      console.log("userInfo", userInfo);
      console.log("accessToken", accessToken);
      console.log("sub", sub);
      throw new Error("email and name are required");
    }

    sessionUserData.set("loginId", userInfo.email);
    sessionUserData.set("name", userInfo.name);

    return redirect("/", {
      headers: [
        ["Set-Cookie", await destroyLoginDataSession(loginDataSession)],
        ["Set-Cookie", await commitUserDataSession(sessionUserData)],
      ],
    });
  } catch (_e: unknown) {
    throw redirect("/login", {
      headers: {
        "Set-Cookie": await destroyLoginDataSession(loginDataSession),
      },
    });
  }
}

export default function Callback() {
  return <div>Index</div>;
}
