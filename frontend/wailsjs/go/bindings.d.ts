import * as models from './models';

export interface go {
  "main": {
    "App": {
		GetNodeData(arg1:string,arg2:string,arg3:number):Promise<string>
		LoadingConnInfo(arg1:string):Promise<string>
		LoadingConnKey():Promise<string>
		LoadingDbResource(arg1:string):Promise<string>
		Ok(arg1:string):Promise<string>
		TestConnection(arg1:string):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
