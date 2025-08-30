import type { LoaderFunctionArgs } from "react-router";
import { Outlet } from "react-router";
import { AppSidebar } from "~/components/app-sidebar";
import { SidebarProvider, SidebarTrigger } from "~/components/ui/sidebar";
import { requireAuthUser } from "~/session.server";

export async function loader({ request }: LoaderFunctionArgs) {
  console.log("_main.tsx::loader");
  const authUser = await requireAuthUser(request);
  return new Response(null, {
    headers: authUser.headers,
  });
}

export default function Layout() {
  return (
    <div>
      <SidebarProvider>
        <AppSidebar />
        <main className="flex-1 overflow-y-scroll px-8 pt-8 ">
          <SidebarTrigger />
          <Outlet />
        </main>
      </SidebarProvider>
    </div>
  );
}
