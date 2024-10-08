# python knn_script.py -x 1.5 -y 2.5 -k 3
import argparse
import math


def euclidean_distance(point1, point2):
    return math.sqrt((point1[0] - point2[0]) ** 2 + (point1[1] - point2[1]) ** 2)


def knn(points, query_point, k):
    distances = [(point, euclidean_distance(point, query_point)) for point in points]
    distances.sort(key=lambda x: x[1])
    return distances[:k]


def main():
    parser = argparse.ArgumentParser(description="Calculate KNN from a text file.")
    parser.add_argument(
        "-x", type=float, required=True, help="X coordinate of the query point"
    )
    parser.add_argument(
        "-y", type=float, required=True, help="Y coordinate of the query point"
    )
    parser.add_argument(
        "-k", type=int, required=True, help="Number of nearest neighbors to find"
    )

    args = parser.parse_args()

    # Read points from file
    points = []
    with open("data.txt", "r") as f:
        for line in f:
            x, y = map(float, line.strip().split())
            points.append((x, y))

    query_point = (args.x, args.y)

    nearest_neighbors = knn(points, query_point, args.k)

    print(f"The {args.k} nearest neighbors to point ({args.x}, {args.y}) are:")
    for i, (point, distance) in enumerate(nearest_neighbors, 1):
        print(f"{i}. ({point[0]:.4f}, {point[1]:.4f}) - Distance: {distance:.4f}")


if __name__ == "__main__":
    main()
