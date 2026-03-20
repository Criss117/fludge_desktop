import { Link, type LinkProps } from "@tanstack/react-router";
import {
  Plus,
  LayoutDashboard,
  Package,
  Users,
  CirclePile,
  Container,
  type LucideIcon,
} from "lucide-react";

import {
  SidebarGroup,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
} from "@shared/components//ui/sidebar";

type NavMenuItem = {
  title: string;
  url: NonNullable<LinkProps["to"]>;
  icon: LucideIcon;
  elements?: NavMenuItem[];
};

interface GroupItemProps extends Props {
  item: NavMenuItem;
}

interface SidebarMenuLinkProps {
  item: NavMenuItem;
  orgId: string;
}

interface Props {
  orgId: string;
}

const navMainItems: NavMenuItem[] = [
  {
    title: "Inicio",
    url: "/dashboard/$orgid",
    icon: LayoutDashboard,
  },
  {
    title: "Inventario",
    url: "/dashboard/$orgid/inventory",
    icon: CirclePile,
    elements: [
      {
        title: "Productos",
        url: "/dashboard/$orgid/inventory/products",
        icon: Package,
      },
      {
        title: "Proveedores",
        url: "/dashboard/$orgid/inventory/suppliers",
        icon: Container,
      },
    ],
  },
  {
    title: "Clientes",
    url: "/dashboard/$orgid/clients",
    icon: Users,
  },
  // {
  //   title: "Equipos",
  //   url: "/dashboard/$orgslug/teams",
  //   icon: Briefcase,
  // },
  // {
  //   title: "Empleados",
  //   url: "/dashboard/$orgslug/employees",
  //   icon: UserCog,
  // },
];

function SidebarMenuLink({ item, orgId }: SidebarMenuLinkProps) {
  return (
    <SidebarMenuButton
      tooltip={item.title}
      render={(props) => {
        return (
          <Link
            to={item.url}
            params={{ orgid: orgId }}
            activeProps={{
              className: "border",
            }}
            inactiveProps={{
              className: "border border-transparent",
            }}
            activeOptions={{
              exact: true,
            }}
            {...props}
          >
            <item.icon />
            <span>{item.title}</span>
          </Link>
        );
      }}
    />
  );
}

function GroupItem({ orgId, item }: GroupItemProps) {
  return (
    <SidebarMenuItem className="space-y-2">
      <SidebarMenuLink item={item} orgId={orgId} />

      <SidebarMenuSub>
        {item.elements?.map((subItem) => (
          <SidebarMenuSubItem key={subItem.title}>
            <SidebarMenuSubButton
              render={(props) => (
                <Link
                  to={subItem.url}
                  params={{ orgid: orgId }}
                  activeProps={{
                    className: "border",
                  }}
                  inactiveProps={{
                    className: "border border-transparent",
                  }}
                  activeOptions={{
                    exact: true,
                  }}
                  {...props}
                >
                  <subItem.icon />
                  <span>{subItem.title}</span>
                </Link>
              )}
            />
          </SidebarMenuSubItem>
        ))}
      </SidebarMenuSub>
    </SidebarMenuItem>
  );
}

export function NavMain({ orgId }: Props) {
  return (
    <SidebarGroup>
      <SidebarGroupContent className="flex flex-col gap-6">
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton
              tooltip="Quick Create"
              className="bg-primary text-primary-foreground hover:bg-primary/90 active:bg-primary/80 min-w-8 transition-colors"
            >
              <Plus />
              <span>Quick Create</span>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>

        <SidebarMenu className="space-y-2">
          {navMainItems.map((item) => {
            if (item.elements)
              return <GroupItem key={item.title} orgId={orgId} item={item} />;

            return (
              <SidebarMenuItem key={item.title}>
                <SidebarMenuLink item={item} orgId={orgId} />
              </SidebarMenuItem>
            );
          })}
        </SidebarMenu>
      </SidebarGroupContent>
    </SidebarGroup>
  );
}
