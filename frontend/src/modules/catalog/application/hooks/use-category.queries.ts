import { count, ilike, useLiveSuspenseQuery } from "@tanstack/react-db";
import { useCategoryCollection } from "./use-category-collection";

type Filters = {
  name?: string;
};

export function useFindManyCategories(filters?: Filters) {
  const categoriesCollection = useCategoryCollection();

  const name = filters?.name;

  const { data } = useLiveSuspenseQuery(
    (q) => {
      let query = q.from({ category: categoriesCollection });

      if (name)
        query = query.where(({ category }) =>
          ilike(category.name, `%${name}%`),
        );

      return query.orderBy(({ category }) => category.createdAt, "desc");
    },
    [name],
  );

  return data;
}

export function useCountAllCategories() {
  const categoriesCollection = useCategoryCollection();

  const { data } = useLiveSuspenseQuery(
    (q) =>
      q
        .from({ category: categoriesCollection })
        .select(({ category }) => ({ total: count(category.id) }))
        .findOne(),
    [],
  );

  return data?.total ?? 0;
}
