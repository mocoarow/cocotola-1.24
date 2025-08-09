import type { LoaderFunctionArgs, MetaFunction } from "react-router";
import { useLoaderData } from "react-router";

import { requireAuthUser } from "~/session.server";

export const meta: MetaFunction = () => {
  return [{ title: "New Remix App" }];
};

export async function loader({ request }: LoaderFunctionArgs) {
  console.log("_main._index.tsx::loader");
  const userInfo = await requireAuthUser(request);
  return { userInfo };
}

export default function Index() {
  console.log("_main.users.tsx::Index");
  const { userInfo } = useLoaderData<typeof loader>();
  const name = userInfo.name;
  return (
    <div>
      <h1 className="bg-blue-400">Welcome to Admin.</h1>
      {name}
      {userInfo.loginId}
    </div>
  );
}
