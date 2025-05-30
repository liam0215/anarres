#include <qpl/qpl.h>
#include <stdint.h>

void setup_qpl_job(qpl_job *job, uint8_t *src, uint8_t *dst, uint32_t src_size,
		   uint32_t dst_size, qpl_operation op)
{
    // We need to explicitly declare that only the current numa node's
    // accelerator should be used: qpl_job *qpl_job_ptr; job->numa_id = -2; If
    // we want to panic on faults, we need to make sure we configure the
    // accelerator so that Block on Fault is set to 0.
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

uint32_t get_safe_compression_buffer_size(uint32_t src_size)
{
    return qpl_get_safe_deflate_compression_buffer_size(src_size);
}

int compress_software(const uint8_t *src, uint8_t *dst, uint32_t src_size,
		      uint32_t dst_size)
{
    qpl_status status;
    qpl_job *job;
    qpl_operation op = qpl_op_compress;
    qpl_path_t path = qpl_path_software;
    uint32_t job_size = 0;

    status = qpl_get_job_size(path, &job_size);
    if (status != QPL_STS_OK) {
        return -1;
    }

    status = qpl_init_job(sw, job);
    if (status != QPL_STS_OK) {
        return -1;
    }

    setup_qpl_job(job, src, dst, src_size, dst_size, op);
    status = qpl_execute_job(job);
    if (status != QPL_STS_OK) {
        return -1;
    }
    // What to do on error?
    status = qpl_fini_job(job);
    return status != QPL_STS_OK;
}
