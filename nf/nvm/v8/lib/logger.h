#ifndef _NEBULAS_NF_NVM_V8_LOGGER_H_
#define _NEBULAS_NF_NVM_V8_LOGGER_H_

void LogInfof(const char *format, ...);
void LogErrorf(const char *format, ...);
void LogDebugf(const char *format, ...);
void LogWarnf(const char *format, ...);
void LogFatalf(const char *format, ...);

#endif // _NEBULAS_NF_NVM_V8_LOGGER_H_
