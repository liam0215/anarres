#!/usr/bin/env -S uv run --script
# /// script
# dependencies = ["matplotlib", "numpy"]
# ///
import csv
import argparse
import matplotlib.pyplot as plt

def read_latency_file(filepath):
    """
    Reads a latency CSV file and returns a list of dicts with keys:
      - param (e.g., "size_kb" or "tput")
      - value (int)
      - median_us (float)
      - p90_us (float)
    """
    data = []
    with open(filepath) as f:
        reader = csv.DictReader(f)
        for row in reader:
            data.append({
                "param": row["param"],
                "value": int(row["value"]),
                "p50_us": float(row["median_us"]),
                "p90_us": float(row["p90_us"]),
            })
    return data


def filter_and_sort(data, x_axis):
    """
    Filters the data by the given x_axis param and sorts by its numeric value.
    x_axis: "size" or "tput"
    """
    key = "size_kb" if x_axis == "size" else "tput"
    filtered = [d for d in data if d["param"] == key]
    return sorted(filtered, key=lambda d: d["value"])


def plot_latency_compare(data1, label1, data2, label2, x_axis, y_axis):
    """
    Plots two latency series on one graph with log-scaled axes.
    y_axis: "median" or "p90" (uses the corresponding _us value)
    """
    x_label = "Size (KB)" if x_axis == "size" else "Offered Load (tput)"
    y_label = f"{y_axis.upper()} Latency (µs)"
    title = f"{y_axis.upper()} Latency vs {x_label} ({label1} vs {label2})"

    series1 = filter_and_sort(data1, x_axis)
    series2 = filter_and_sort(data2, x_axis)

    x1 = [d["value"] for d in series1]
    y1 = [d[f"{y_axis}_us"] for d in series1]

    x2 = [d["value"] for d in series2]
    y2 = [d[f"{y_axis}_us"] for d in series2]

    plt.figure(figsize=(8, 5))
    plt.plot(x1, y1, marker='o', label=label1)
    plt.plot(x2, y2, marker='s', label=label2)
    plt.xlabel(x_label)
    plt.ylabel(y_label)
    plt.title(title)
    plt.grid(True)
    plt.xscale('log')  # Set X axis to log scale
    plt.yscale('log')  # Set Y axis to log scale
    plt.legend()
    plt.tight_layout()

    output_file = f"{y_axis}_vs_{x_axis}_compare.png"
    plt.savefig(output_file)
    plt.show()


if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Compare latency metrics between two datasets with log-scaled axes"
    )
    parser.add_argument("csv1", help="Path to first latency CSV file (e.g. software)")
    parser.add_argument("csv2", help="Path to second latency CSV file (e.g. hardware)")
    parser.add_argument("--label1", default="SW", help="Label for first dataset")
    parser.add_argument("--label2", default="HW", help="Label for second dataset")
    parser.add_argument(
        "--x", choices=["size", "tput"], required=True,
        help="X-axis: size or tput"
    )
    parser.add_argument(
        "--y", choices=["p50", "p90"], required=True,
        help="Y-axis: p50 or p90"
    )
    args = parser.parse_args()

    data1 = read_latency_file(args.csv1)
    data2 = read_latency_file(args.csv2)

    plot_latency_compare(
        data1, args.label1,
        data2, args.label2,
        args.x, args.y
    )
