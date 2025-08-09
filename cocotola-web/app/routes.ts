import type { RouteConfig } from "@react-router/dev/routes";
import { remixRoutesOptionAdapter } from "@react-router/remix-routes-option-adapter";

export default remixRoutesOptionAdapter((defineRoutes) => {
  return defineRoutes((route) => {
    // 認証系
    route("login", "./routes/_auth.login.tsx");
    route("callback", "./routes/_auth.callback.tsx");
    route("logout", "./routes/_auth.logout.tsx");

    // メインレイアウト
    route("", "./routes/_main.tsx", () => {
      // index
      route("", "./routes/_main._index.tsx", { index: true });
    });
  });
}) satisfies RouteConfig;
