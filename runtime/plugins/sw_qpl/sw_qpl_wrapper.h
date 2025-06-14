#ifndef QPL_WRAPPER_H
#define QPL_WRAPPER_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stddef.h>
#include <stdint.h>

#include "qpl/qpl.h"

qpl_status sw_qpl_compress_wrapper(
    uint8_t *src,
    uint32_t src_len,
    uint8_t **dst,
    uint32_t *dst_len,
    qpl_path_t execution_path);

qpl_status sw_qpl_decompress_wrapper(uint8_t *src,
                                  uint32_t src_len,
                                  uint32_t expected_len,
                                  uint8_t **dst,
                                  uint32_t *dst_len,
                                  qpl_path_t execution_path);

#ifdef __cplusplus
}
#endif /* extern "C" */

#endif /* QPL_WRAPPER_H */
