export namespace dtos {
	
	export class SummaryOperatorDTO {
	    id: string;
	    name: string;
	    username: string;
	    email: string;
	    isRoot: boolean;
	    createdAt: number;
	    updatedAt: number;
	    deletedAt?: number;
	
	    static createFrom(source: any = {}) {
	        return new SummaryOperatorDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.isRoot = source["isRoot"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.deletedAt = source["deletedAt"];
	    }
	}
	export class AppStateDTO {
	    activeOrganizationId?: string;
	    activeOperator?: SummaryOperatorDTO;
	    updatedAt: number;
	    operators: SummaryOperatorDTO[];
	
	    static createFrom(source: any = {}) {
	        return new AppStateDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activeOrganizationId = source["activeOrganizationId"];
	        this.activeOperator = this.convertValues(source["activeOperator"], SummaryOperatorDTO);
	        this.updatedAt = source["updatedAt"];
	        this.operators = this.convertValues(source["operators"], SummaryOperatorDTO);
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
	export class SignInDTO {
	    username: string;
	    pin: string;
	
	    static createFrom(source: any = {}) {
	        return new SignInDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.pin = source["pin"];
	    }
	}
	export class SignupDTO {
	    name: string;
	    username: string;
	    email: string;
	    pin: string;
	
	    static createFrom(source: any = {}) {
	        return new SignupDTO(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.username = source["username"];
	        this.email = source["email"];
	        this.pin = source["pin"];
	    }
	}

}

