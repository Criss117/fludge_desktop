import type { Category } from "@catalog/domain/entities/category.entity";
import type {
  CategoryRepository,
  CreateCategoryCommand,
  DelelteManyCategoriesCommand,
  UpdateCategoryCommand,
} from "@catalog/domain/ports/category.repository";

export class CategoryService {
  private readonly categoryRepository: CategoryRepository;

  constructor(categoryRepository: CategoryRepository) {
    this.categoryRepository = categoryRepository;
  }

  public async createCategory(cmd: CreateCategoryCommand): Promise<Category> {
    return this.categoryRepository.craeteCategory(cmd);
  }

  public async deleteManyCategories(
    cmd: DelelteManyCategoriesCommand,
  ): Promise<void> {
    return this.categoryRepository.deleteManyCategories(cmd);
  }

  public async updateCategory(cmd: UpdateCategoryCommand): Promise<Category> {
    return this.categoryRepository.updateCategory(cmd);
  }

  public async findAllCategories(): Promise<Category[]> {
    return this.categoryRepository.findAllCategories();
  }
}
