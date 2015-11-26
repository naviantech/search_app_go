package search

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

type Search struct {
	key         string
	results     map[string]string
	num_results int
}

func (s *Search) Results() map[string]string {
	return s.results
}
