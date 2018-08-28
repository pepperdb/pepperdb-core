#ifndef _NEBULAS_NF_NVM_V8_LIB_FAKE_BLOCKCHAIN_H_
#define _NEBULAS_NF_NVM_V8_LIB_FAKE_BLOCKCHAIN_H_

#include <stddef.h>

char *GetTxByHash(void *handler, const char *hash, size_t *gasCnt);
int GetAccountState(void *handler, const char *addres, size_t *gasCnts, char **result, char **info);
int Transfer(void *handler, const char *to, const char *value, size_t *gasCnt);
int VerifyAddress(void *handler, const char *address, size_t *gasCnt);
int GetPreBlockHash(void *handler, unsigned long long offset, size_t *counterVal, char **result, char **info);
int GetPreBlockSeed(void *handler, unsigned long long offset, size_t *counterVal, char **result, char **info);


#endif //_NEBULAS_NF_NVM_V8_LIB_FAKE_BLOCKCHAIN_H_
