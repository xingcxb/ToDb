export interface go {
  "main": {
    "App": {
		TestConnection(arg1:string):Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
