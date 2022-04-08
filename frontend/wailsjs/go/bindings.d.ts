import * as models from './models';

export interface go {
  "main": {
    "App": {
		GetNodeData(connType:string,connName:string,nodeId: number): Promise<string>
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
