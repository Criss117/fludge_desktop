import { Link } from "@tanstack/react-router";
import { Check, ChevronDown, Building2, PlusCircleIcon } from "lucide-react";

import {
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@shared/components/ui/sidebar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@shared/components/ui/dropdown-menu";
import { Logo } from "@shared/components/logo";
import { cn } from "@shared/lib/utils";

import { useIam } from "@/integrations/iam";

interface Props {
  orgId: string;
}

export function AppSideBarHeader({ orgId }: Props) {
  const { appState } = useIam();

  const activeOrganization = appState?.activeOrganization;
  const activeOperator = appState?.activeOperator;

  return (
    <SidebarHeader>
      <SidebarMenu>
        <SidebarMenuItem className="flex items-center">
          <SidebarMenuButton
            className="[&_svg]:size-8"
            render={(props) => (
              <Link
                to="/dashboard/$orgid"
                params={{ orgid: orgId }}
                {...props}
              />
            )}
          >
            <Logo />
            <span className="text-base font-semibold">Fludge</span>
          </SidebarMenuButton>
        </SidebarMenuItem>

        <SidebarMenuItem>
          <DropdownMenu>
            <DropdownMenuTrigger
              render={
                <SidebarMenuButton className="cursor-pointer justify-between">
                  <div className="flex items-center gap-2">
                    <Building2 className="size-4" />
                    <span className="truncate font-medium">
                      {activeOrganization?.name || "Organización"}
                    </span>
                  </div>
                  <ChevronDown className="size-4 text-muted-foreground" />
                </SidebarMenuButton>
              }
            />
            <DropdownMenuContent side="right" align="start" className="w-64">
              <DropdownMenuGroup>
                <DropdownMenuLabel>Organizaciones</DropdownMenuLabel>
                <DropdownMenuSeparator />
                {activeOperator?.isRoot &&
                  activeOperator?.isMemberIn.map((org) => {
                    const isActive = org.id === activeOrganization?.id;

                    return (
                      <DropdownMenuItem
                        key={org.id}
                        className="cursor-pointer flex items-center justify-between"
                        render={(props) => (
                          <Link
                            to="/dashboard/$orgid"
                            preload={false}
                            params={{ orgid: org.id }}
                            {...props}
                          />
                        )}
                      >
                        <div className="flex items-center gap-2">
                          <Building2 className="size-4 text-muted-foreground" />
                          <span className="truncate">{org.name}</span>
                        </div>
                        <Check
                          className={cn(
                            "size-4",
                            isActive ? "opacity-100" : "opacity-0",
                          )}
                        />
                      </DropdownMenuItem>
                    );
                  })}
              </DropdownMenuGroup>
              <DropdownMenuSeparator />
              <DropdownMenuGroup>
                <DropdownMenuItem
                  nativeButton={false}
                  className="cursor-pointer flex items-center"
                  render={(props) => (
                    <Link {...props} to="/register-organization" />
                  )}
                >
                  <PlusCircleIcon />
                  <span>Registrar Organización</span>
                </DropdownMenuItem>
              </DropdownMenuGroup>
            </DropdownMenuContent>
          </DropdownMenu>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarHeader>
  );
}
