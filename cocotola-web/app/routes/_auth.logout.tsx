import type { MetaFunction } from "react-router";
import { type LoaderFunctionArgs, redirect } from "react-router";

import { destroyUserDataSession, getUserDataSession } from "~/session.server";

export const meta: MetaFunction = () => {
  return [{ title: "Logout" }];
};

export async function loader({ request }: LoaderFunctionArgs) {
  console.log("_auth.logout.tsx::loader");

  const sessionUserData = await getUserDataSession(request);
  return redirect("/login", {
    headers: {
      "Set-Cookie": await destroyUserDataSession(sessionUserData),
    },
  });
}

export default function Logout() {
  return <div />;
}
