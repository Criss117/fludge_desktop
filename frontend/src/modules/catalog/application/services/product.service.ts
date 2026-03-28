import type { Product } from "@catalog/domain/entities/product.entity";
import type {
  CreateProductCommand,
  ProductRepository,
  UpdateProductCommand,
} from "@catalog/domain/ports/product.repository";

export class ProductService {
  private readonly productRepository: ProductRepository;

  constructor(productRepository: ProductRepository) {
    this.productRepository = productRepository;
  }

  public async findAllProducts(): Promise<Product[]> {
    return this.productRepository.findAllProducts();
  }

  public async createProduct(cmd: CreateProductCommand): Promise<Product> {
    return this.productRepository.createProduct(cmd);
  }

  public async updateProduct(cmd: UpdateProductCommand): Promise<Product> {
    return this.productRepository.updateProduct(cmd);
  }
}
