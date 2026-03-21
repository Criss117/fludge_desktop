import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function slugify(text: string): string {
  return (
    text
      .toLowerCase()
      .trim()
      // Reemplazar caracteres acentuados
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "")
      // Reemplazar espacios y guiones bajos con guiones
      .replace(/[\s_]+/g, "-")
      // Eliminar caracteres especiales excepto guiones
      .replace(/[^\w-]+/g, "")
      // Reemplazar múltiples guiones con uno solo
      .replace(/--+/g, "-")
      // Eliminar guiones al inicio y final
      .replace(/^-+|-+$/g, "")
  );
}

export const inputNumberHelper = {
  value: (value: string, isTouched: boolean) =>
    value.toString() === "0" && !isTouched ? "" : value.toString(),
  onChange: (value: string) => {
    const val = value === "" ? undefined : Number(value);
    return val?.toString() || "";
  },
};

export function formatCurrency(value: number) {
  return new Intl.NumberFormat("es-ES", {
    style: "currency",
    currency: "COP",
  }).format(value);
}
