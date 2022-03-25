export interface go {
  "main": {
    "App": {
		LoadingConnInfo(arg1:string):Promise<string>
		LoadingConnKey():Promise<string>
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
