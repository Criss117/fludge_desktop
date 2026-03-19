import { useRouter } from "@tanstack/react-router";
import { UserCog, Settings, LogOut } from "lucide-react";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@shared/components/ui/dropdown-menu";
import {
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@shared/components/ui/sidebar";
import { Avatar, AvatarFallback } from "@shared/components/ui/avatar";

import { useIam } from "@/integrations/iam";

function getInitials(name: string): string {
  return name
    .split(" ")
    .map((part) => part[0])
    .join("")
    .toUpperCase()
    .slice(0, 2);
}

export function NavUser() {
  const router = useRouter();
  const { signOut, appState } = useIam();

  const activeOperator = appState?.activeOperator;

  if (!activeOperator) return null;

  const handleSignOut = () => {
    signOut.mutateAsync().then(() => {
      router.navigate({ to: "/" });
    });
  };

  return (
    <SidebarMenu>
      <SidebarMenuItem>
        <DropdownMenu>
          <DropdownMenuTrigger
            render={
              <SidebarMenuButton className="cursor-pointer">
                <Avatar className="size-8">
                  {/* {session.user.image && (
                    <AvatarImage
                      src={session.user.image}
                      alt={session.user.name}
                    />
                  )} */}
                  <AvatarFallback className="text-xs">
                    {activeOperator.name
                      ? getInitials(activeOperator.name)
                      : "U"}
                  </AvatarFallback>
                </Avatar>
                <span className="truncate text-sm font-medium">
                  {activeOperator.name || "User"}
                </span>
              </SidebarMenuButton>
            }
          />
          <DropdownMenuContent side="left" align="end" className="w-56">
            <DropdownMenuGroup>
              <DropdownMenuLabel>
                <div className="flex flex-col space-y-1">
                  <span className="truncate font-medium">
                    {activeOperator.name}
                  </span>
                  <span className="truncate text-xs text-muted-foreground">
                    {activeOperator.email}
                  </span>
                </div>
              </DropdownMenuLabel>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem>
              <UserCog className="mr-2 size-4" />
              <span>Perfil</span>
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Settings className="mr-2 size-4" />
              <span>Configuración</span>
            </DropdownMenuItem>
            <DropdownMenuSeparator />
            <DropdownMenuItem
              className="text-destructive focus:text-destructive"
              onClick={handleSignOut}
            >
              <LogOut className="mr-2 size-4" />
              <span>Cerrar sesión</span>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </SidebarMenuItem>
    </SidebarMenu>
  );
}
