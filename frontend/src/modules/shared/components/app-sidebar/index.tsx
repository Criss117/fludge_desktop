import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarSeparator,
} from "../ui/sidebar";
import { NavSystem } from "./nav-system";
import { NavMain } from "./nav-main";
import { NavUser } from "./nav-user";
import { AppSideBarHeader } from "./header";

export function AppSidebar({ orgId }: { orgId: string }) {
  return (
    <Sidebar variant="inset" collapsible="icon">
      <AppSideBarHeader orgId={orgId} />
      <SidebarContent>
        <NavMain orgId={orgId} />
        <SidebarSeparator className="my-2" />
        <NavSystem />
      </SidebarContent>
      <SidebarFooter>
        <NavUser />
      </SidebarFooter>
    </Sidebar>
  );
}
