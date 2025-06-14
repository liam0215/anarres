#!/usr/bin/env -S uv run --script
# /// script
# dependencies = ["matplotlib", "numpy"]
# ///
import subprocess
import csv
import time
from pathlib import Path
import statistics
import numpy as np
import argparse

# Paths
workload_script = "./run_workload.sh"
stats_file = Path("build/cmplx_wlgen/cmplx_wlgen_proc/cmplx_wlgen_proc/stats.csv")

def run_workload(size=None, tput=None):
    args = []
    if size is not None:
        args.append(str(size))
    else:
        args.append("64")
    if tput is not None:
        args.append(str(tput))
    subprocess.run([workload_script] + args, check=True)
    time.sleep(1)

def parse_stats():
    if not stats_file.exists():
        print("Warning: stats.csv not found")
        return []

    with open(stats_file) as f:
        reader = csv.DictReader(f)
        return [
            int(row["Duration"]) / 1000  # ns → us
            for row in reader if row["IsError"].lower() == "false"
        ]

def write_header(outfile):
    with open(outfile, "w", newline="") as f:
        writer = csv.writer(f)
        writer.writerow([
            "param", "value", "min_us", "max_us", "avg_us",
            "median_us", "p90_us", "p99_us", "num_samples"
        ])

def append_result(outfile, param_name, param_val, latencies):
    latencies = np.array(latencies)
    row = [
        param_name,
        param_val,
        round(latencies.min(), 2),
        round(latencies.max(), 2),
        round(latencies.mean(), 2),
        round(np.median(latencies), 2),
        round(np.percentile(latencies, 90), 2),
        round(np.percentile(latencies, 99), 2),
        len(latencies)
    ]
    with open(outfile, "a", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(row)

def sweep_size():
    outfile = "latency_results_size.csv"
    write_header(outfile)

    for size_kb in [2 ** i for i in range(0, 9)]:  # 1 → 256
        print(f"Running workload for --size={size_kb} KB")
        stats_file.unlink(missing_ok=True)
        run_workload(size=size_kb)
        latencies = parse_stats()
        if latencies:
            append_result(outfile, "size_kb", size_kb, latencies)
        else:
            print(f"Warning: No data for size {size_kb}")

def sweep_tput():
    outfile = "latency_results_tput.csv"
    write_header(outfile)

    for tput in range(1000, 8001, 1000):
        print(f"Running workload for --tput={tput}")
        stats_file.unlink(missing_ok=True)
        run_workload(tput=tput)
        latencies = parse_stats()
        if latencies:
            append_result(outfile, "tput", tput, latencies)
        else:
            print(f"Warning: No data for tput {tput}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--sweep-size", action="store_true", help="Sweep --size from 1KB to 256KB")
    parser.add_argument("--sweep-tput", action="store_true", help="Sweep --tput from 1000 to 8000")
    args = parser.parse_args()

    if args.sweep_size:
        sweep_size()
    elif args.sweep_tput:
        sweep_tput()
    else:
        print("Please specify --sweep-size or --sweep-tput")
