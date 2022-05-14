import * as models from './models';

export interface go {
  "main": {
    "App": {
		ExportConn():Promise<void>
		GetNodeData(arg1:string,arg2:string,arg3:string):Promise<string>
		ImportConn():Promise<void>
		LoadingConnInfo(arg1:string,arg2:string):Promise<string>
		LoadingConnKey():Promise<string>
		LoadingDbResource(arg1:string):Promise<string>
		Ok(arg1:string):Promise<string>
		RedisDelKey(arg1:string,arg2:string,arg3:string,arg4:string):Promise<void>
		RedisGetData(arg1:string,arg2:string,arg3:string,arg4:string):Promise<string>
		RedisReName(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<void>
		RedisSaveStringValue(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string,arg6:string):Promise<void>
		RedisUpTtl(arg1:string,arg2:string,arg3:string,arg4:string,arg5:string):Promise<void>
		TestConnection(arg1:string):Promise<void>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
