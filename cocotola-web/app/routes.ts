import { index, type RouteConfig, route } from "@react-router/dev/routes";
export default [
  route("login", "./routes/_auth.login.tsx"),
  route("callback", "./routes/_auth.callback.tsx"),
  route("logout", "./routes/_auth.logout.tsx"),

  route("", "./routes/_main.tsx", [index("./routes/_main._index.tsx")]),
] satisfies RouteConfig;
