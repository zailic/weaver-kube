// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package impl

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// greenText returns the ANSI escape code for a green colored text.
func greenText() string {
	return "\033[32m%s\033[0m\n"
}

// cp copies the src file to the dst files.
//
// TODO(rgrandl): remove duplicate.
func cp(src, dst string) error {
	// Open src.
	srcf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("open %q: %w", src, err)
	}
	defer srcf.Close()
	srcinfo, err := srcf.Stat()
	if err != nil {
		return fmt.Errorf("stat %q: %w", src, err)
	}

	// Create or truncate dst.
	dstf, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("create %q: %w", dst, err)
	}
	defer dstf.Close()

	// Copy src to dst.
	const bufSize = 1 << 20
	if _, err := io.Copy(dstf, bufio.NewReaderSize(srcf, bufSize)); err != nil {
		return fmt.Errorf("cp %q %q: %w", src, dst, err)
	}
	if err := os.Chmod(dst, srcinfo.Mode()); err != nil {
		return fmt.Errorf("chmod %q: %w", dst, err)
	}
	return nil
}

// TODO(rgrandl): Remove duplicate.
func ptrOf[T any](val T) *T { return &val }

// topologicalSort returns a topological sort of the given graph.
// The algorithm is based on https://en.wikipedia.org/wiki/Topological_sorting#Depth-first_search,
// and is not optimized for performance, but it is simple and works for our use case.
// Note that the algorithm is not guaranteed to return a consistent result if the input graph has cycles.
func topologicalSort(edges [][2]string) []string {
	// transform edges into a map from node to its dependencies
	m := make(map[string][]string)
	for _, e := range edges {
		m[e[0]] = append(m[e[0]], e[1])
	}

	var sorted []string
	visited := make(map[string]bool)
	// visit is a recursive function that visits a node and its dependencies
	var visit func(string)
	visit = func(node string) {
		if visited[node] {
			return
		}
		visited[node] = true
		for _, dep := range m[node] {
			visit(dep)
		}
		sorted = append(sorted, node)
	}
	for node := range m {
		visit(node)
	}

	return sorted
}
