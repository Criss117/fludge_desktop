export namespace appstate {
	
	export class ResponseSessionState {
	    activeOrganization?: responses.OrganizationResponse;
	    activeOperator?: responses.OperatorResponse;
	
	    static createFrom(source: any = {}) {
	        return new ResponseSessionState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activeOrganization = this.convertValues(source["activeOrganization"], responses.OrganizationResponse);
	        this.activeOperator = this.convertValues(source["activeOperator"], responses.OperatorResponse);
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

export namespace commands {
	
	export class SignInCommand {
	    username: string;
	    pin: string;
	
	    static createFrom(source: any = {}) {
	        return new SignInCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.username = source["username"];
	        this.pin = source["pin"];
	    }
	}
	export class SignUpCommand {
	    name: string;
	    email: string;
	    username: string;
	    pin: string;
	
	    static createFrom(source: any = {}) {
	        return new SignUpCommand(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.email = source["email"];
	        this.username = source["username"];
	        this.pin = source["pin"];
	    }
	}

}

export namespace responses {
	
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
	export class SignInResponse {
	    activeOperator: OperatorResponse;
	    activeTeams: TeamResponse[];
	
	    static createFrom(source: any = {}) {
	        return new SignInResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.activeOperator = this.convertValues(source["activeOperator"], OperatorResponse);
	        this.activeTeams = this.convertValues(source["activeTeams"], TeamResponse);
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

