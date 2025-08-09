import { Outlet } from "react-router";

import { AppSidebar } from "~/components/app-sidebar";
import { SidebarProvider, SidebarTrigger } from "~/components/ui/sidebar";

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
