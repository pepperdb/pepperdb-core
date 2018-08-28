#ifndef _NEBULAS_NF_NVM_V8_ERROR_H_
#define _NEBULAS_NF_NVM_V8_ERROR_H_

/*
success or crash
*/
#define REPORT_UNEXPECTED_ERR() do{ \
  if (NULL == isolate) {\
    LogFatalf("Unexpected Error: invalid argument, ioslate is NULL");\
  }\
  Local<Context> context = isolate->GetCurrentContext();\
  V8Engine *e = GetV8EngineInstance(context);\
  if (NULL == e) {\
    LogFatalf("Unexpected Error: failed to get V8Engine");\
  }\
  TerminateExecution(e);\
  e->is_unexpected_error_happen = true;\
} while(0)

#define DEAL_ERROR_FROM_GOLANG(err) do {\
  if (NVM_UNEXPECTED_ERR == err || (NVM_EXCEPTION_ERR == err && NULL == exceptionInfo) ||\
    (NVM_SUCCESS == err && NULL == result)) {\
    info.GetReturnValue().SetNull();\
    REPORT_UNEXPECTED_ERR();\
  } else if (NVM_EXCEPTION_ERR == err) {\
    isolate->ThrowException(String::NewFromUtf8(isolate, exceptionInfo));\
  } else if (NVM_SUCCESS == err) {\
    info.GetReturnValue().Set(String::NewFromUtf8(isolate, result));\
  } else {\
    info.GetReturnValue().SetNull();\
    REPORT_UNEXPECTED_ERR();\
  }\
} while(0)

enum nvmErrno {
  NVM_SUCCESS = 0,
  NVM_EXCEPTION_ERR = -1,
  NVM_MEM_LIMIT_ERR = -2,
  NVM_GAS_LIMIT_ERR = -3,
  NVM_UNEXPECTED_ERR = -4,
  NVM_EXE_TIMEOUT_ERR = -5,
};

#endif //_NEBULAS_NF_NVM_V8_ENGINE_ERROR_H_