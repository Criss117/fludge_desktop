export namespace commands {
	
	export class CreateCategory {
	    name: string;
	    description?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateCategory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class CreateProduct {
	    name: string;
	    sku: string;
	    description?: string;
	    wholesalePrice: number;
	    salePrice: number;
	    costPrice: number;
	    stock: number;
	    minStock: number;
	    categoryId?: string;
	    supplierId?: string;
	
	    static createFrom(source: any = {}) {
	        return new CreateProduct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.sku = source["sku"];
	        this.description = source["description"];
	        this.wholesalePrice = source["wholesalePrice"];
	        this.salePrice = source["salePrice"];
	        this.costPrice = source["costPrice"];
	        this.stock = source["stock"];
	        this.minStock = source["minStock"];
	        this.categoryId = source["categoryId"];
	        this.supplierId = source["supplierId"];
	    }
	}
	export class DeleteManyCategories {
	    ids: string[];
	
	    static createFrom(source: any = {}) {
	        return new DeleteManyCategories(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ids = source["ids"];
	    }
	}
	export class RegisterOrganization {
	    name: string;
	    legalName: string;
	    address: string;
	    logo?: string;
	    contactPhone?: string;
	    contactEmail?: string;
	
	    static createFrom(source: any = {}) {
	        return new RegisterOrganization(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.legalName = source["legalName"];
	        this.address = source["address"];
	        this.logo = source["logo"];
	        this.contactPhone = source["contactPhone"];
	        this.contactEmail = source["contactEmail"];
	    }
	}
	export class SignIn {
	    username: string;
	    pin: string;
	
	    static createFrom(source: any = {}) {
	        return new SignIn(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.pin = source["pin"];
	    }
	}
	export class SignUp {
	    name: string;
	    email: string;
	    username: string;
	    pin: string;
	
	    static createFrom(source: any = {}) {
	        return new SignUp(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.email = source["email"];
	        this.username = source["username"];
	        this.pin = source["pin"];
	    }
	}
	export class SwitchOrganization {
	    organizationId: string;
	
	    static createFrom(source: any = {}) {
	        return new SwitchOrganization(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.organizationId = source["organizationId"];
	    }
	}
	export class UpdateCategory {
	    id: string;
	    name: string;
	    description?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateCategory(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	    }
	}
	export class UpdateProduct {
	    id: string;
	    name: string;
	    sku: string;
	    description?: string;
	    wholesalePrice: number;
	    salePrice: number;
	    costPrice: number;
	    stock: number;
	    minStock: number;
	    categoryId?: string;
	    supplierId?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateProduct(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.sku = source["sku"];
	        this.description = source["description"];
	        this.wholesalePrice = source["wholesalePrice"];
	        this.salePrice = source["salePrice"];
	        this.costPrice = source["costPrice"];
	        this.stock = source["stock"];
	        this.minStock = source["minStock"];
	        this.categoryId = source["categoryId"];
	        this.supplierId = source["supplierId"];
	    }
	}

}

export namespace responses {
	
	export class CategoryResponse {
	    id: string;
	    name: string;
	    description?: string;
	    organizationId: string;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	
	    static createFrom(source: any = {}) {
	        return new CategoryResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.organizationId = source["organizationId"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	    }
	}
	export class MemberResponse {
	    id: string;
	    organizationId: string;
	    operatorId: string;
	    role: string;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	
	    static createFrom(source: any = {}) {
	        return new MemberResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.organizationId = source["organizationId"];
	        this.operatorId = source["operatorId"];
	        this.role = source["role"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	    }
	}
	export class OperatorOrganizationResponse {
	    id: string;
	    slug: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new OperatorOrganizationResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.slug = source["slug"];
	        this.name = source["name"];
	    }
	}
	export class OperatorResponse {
	    id: string;
	    name: string;
	    email: string;
	    username: string;
	    pin: string;
	    isRoot: boolean;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	    isMemberIn: OperatorOrganizationResponse[];
	
	    static createFrom(source: any = {}) {
	        return new OperatorResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.email = source["email"];
	        this.username = source["username"];
	        this.pin = source["pin"];
	        this.isRoot = source["isRoot"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	        this.isMemberIn = this.convertValues(source["isMemberIn"], OperatorOrganizationResponse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class TeamMemberResponse {
	    id: string;
	    teamId: string;
	    operatorId: string;
	    organizationId: string;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	
	    static createFrom(source: any = {}) {
	        return new TeamMemberResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.teamId = source["teamId"];
	        this.operatorId = source["operatorId"];
	        this.organizationId = source["organizationId"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	    }
	}
	export class TeamResponse {
	    id: string;
	    name: string;
	    organizationId: string;
	    permissions: string[];
	    description?: string;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	    Members: TeamMemberResponse[];
	
	    static createFrom(source: any = {}) {
	        return new TeamResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.organizationId = source["organizationId"];
	        this.permissions = source["permissions"];
	        this.description = source["description"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	        this.Members = this.convertValues(source["Members"], TeamMemberResponse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class OrganizationResponse {
	    id: string;
	    name: string;
	    slug: string;
	    logo?: string;
	    metadata: number[];
	    legalName: string;
	    address: string;
	    contactPhone?: string;
	    contactEmail?: string;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	    Members: MemberResponse[];
	    Teams: TeamResponse[];
	
	    static createFrom(source: any = {}) {
	        return new OrganizationResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.slug = source["slug"];
	        this.logo = source["logo"];
	        this.metadata = source["metadata"];
	        this.legalName = source["legalName"];
	        this.address = source["address"];
	        this.contactPhone = source["contactPhone"];
	        this.contactEmail = source["contactEmail"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	        this.Members = this.convertValues(source["Members"], MemberResponse);
	        this.Teams = this.convertValues(source["Teams"], TeamResponse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ProductResponse {
	    id: string;
	    sku: string;
	    name: string;
	    description?: string;
	    wholesalePrice: number;
	    salePrice: number;
	    costPrice: number;
	    stock: number;
	    minStock: number;
	    organizationId: string;
	    categoryId?: string;
	    supplierId?: string;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	
	    static createFrom(source: any = {}) {
	        return new ProductResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.sku = source["sku"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.wholesalePrice = source["wholesalePrice"];
	        this.salePrice = source["salePrice"];
	        this.costPrice = source["costPrice"];
	        this.stock = source["stock"];
	        this.minStock = source["minStock"];
	        this.organizationId = source["organizationId"];
	        this.categoryId = source["categoryId"];
	        this.supplierId = source["supplierId"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	    }
	}
	export class ResponseAppState {
	    activeOrganization?: OrganizationResponse;
	    activeOperator?: OperatorResponse;
	
	    static createFrom(source: any = {}) {
	        return new ResponseAppState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activeOrganization = this.convertValues(source["activeOrganization"], OrganizationResponse);
	        this.activeOperator = this.convertValues(source["activeOperator"], OperatorResponse);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

