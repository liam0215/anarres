/**
 * The following codes are adapted from official examples of Intel QPL library.
 *
 * See https://intel.github.io/qpl/documentation/get_started_docs/compression_decompression.html#compression
 */

#include "sw_qpl_wrapper.h"

#include <stdlib.h>
#include <string.h>

#include <iostream>

void sw_setup_qpl_job(qpl_job *job, uint8_t *src, uint8_t *dst, uint32_t src_size,
		   uint32_t dst_size, qpl_operation op)
{
    // We need to explicitly declare that only the current numa node's
    // accelerator should be used: qpl_job *qpl_job_ptr; job->numa_id = -2; 
    job->numa_id = -2;
    job->op = op;
    job->level = qpl_default_level;
    job->next_in_ptr = src;
    job->next_out_ptr = dst;
    job->available_in = src_size;
    job->available_out = dst_size;
    uint32_t compression_flag =
	(op == qpl_op_compress) * QPL_FLAG_DYNAMIC_HUFFMAN;

    job->flags = QPL_FLAG_FIRST | QPL_FLAG_LAST | QPL_FLAG_OMIT_VERIFY |
		 compression_flag;
}

qpl_status sw_qpl_compress_wrapper(uint8_t *src,
                                uint32_t src_len,
                                uint8_t **dst,
                                uint32_t *dst_len,
                                qpl_path_t execution_path) {
    *dst = NULL;
    *dst_len = 0;

    // Get compression buffer size estimate
    const uint32_t compression_size = qpl_get_safe_deflate_compression_buffer_size(src_len);
    if (compression_size == 0) {
        std::cerr << "Invalid source size. Source size exceeds the maximum supported size." << src_len << "\n";
        return QPL_STS_SIZE_ERR;
    }

    // Job initialization
    uint32_t job_size = 0;
    qpl_status status = qpl_get_job_size(execution_path, &job_size);
    if (status != QPL_STS_OK) {
        std::cerr << "An error " << status << " acquired during job size getting.\n";
        return status;
    }

    void *job_buffer = malloc(job_size);
    qpl_job *job = (qpl_job *)job_buffer;

    status = qpl_init_job(execution_path, job);
    if (status != QPL_STS_OK) {
        std::cerr << "An error " << status << " acquired during job initializing.\n";
        free(job_buffer);
        return status;
    }

    // Performing a compression operation
    uint8_t *out_buf = (uint8_t *)malloc(compression_size);

    sw_setup_qpl_job(job, src, out_buf, src_len, compression_size, qpl_op_compress);

    // Compression
    status = qpl_execute_job(job);
    if (status != QPL_STS_OK) {
        std::cout << "An error " << status << " acquired during compression.\n";
        free(out_buf);
        free(job_buffer);
        return status;
    }

    *dst_len = job->total_out;
    *dst = out_buf;

    // Freeing resources
    status = qpl_fini_job(job);
    if (status != QPL_STS_OK) {
        std::cerr << "An error " << status << " acquired during job finalizing.\n";
        free(out_buf);
        free(job_buffer);
        return status;
    }

    return status;
}

qpl_status sw_qpl_decompress_wrapper(uint8_t *src,
                                  uint32_t src_len,
                                  uint32_t expected_len,
                                  uint8_t **dst,
                                  uint32_t *dst_len,
                                  qpl_path_t execution_path) {
    *dst = NULL;
    *dst_len = 0;

    // Job initialization
    uint32_t job_size = 0;
    qpl_status status = qpl_get_job_size(execution_path, &job_size);
    if (status != QPL_STS_OK) {
        std::cerr << "An error " << status << " acquired during job size getting.\n";
        return status;
    }

    void *job_buffer = malloc(job_size);
    qpl_job *job = (qpl_job *)job_buffer;

    status = qpl_init_job(execution_path, job);
    if (status != QPL_STS_OK) {
        std::cerr << "An error " << status << " acquired during job initializing.\n";
        free(job_buffer);
        return status;
    }

    // Performing a decompression operation
    uint8_t *out_buf = (uint8_t *)malloc(expected_len);

    job->op = qpl_op_decompress;
    job->next_in_ptr = src;
    job->next_out_ptr = out_buf;
    job->available_in = src_len;
    job->available_out = expected_len;
    job->flags = QPL_FLAG_FIRST | QPL_FLAG_LAST;

    // Decompression
    status = qpl_execute_job(job);
    if (status != QPL_STS_OK) {
        std::cout << "An error " << status << " acquired during decompression.\n";
        free(out_buf);
        free(job_buffer);
        return status;
    }

    *dst_len = job->total_out;
    *dst = out_buf;

    // Freeing resources
    status = qpl_fini_job(job);
    if (status != QPL_STS_OK) {
        std::cerr << "An error " << status << " acquired during job finalizing.\n";
        free(out_buf);
        free(job_buffer);
        return status;
    }

    return status;
}
