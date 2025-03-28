export namespace main {
	
	export class SavedGamesPath {
	    dir: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new SavedGamesPath(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dir = source["dir"];
	        this.name = source["name"];
	    }
	}

}

