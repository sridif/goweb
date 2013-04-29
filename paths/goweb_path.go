package paths

import (
	"strings"
)

/*
  GowebPath represents a path that can contain special
  matching segments.

  Valid paths include:

    /literal
    /{placeholder}
    /[optional placeholder]
    /* - matches everything following this
    /people/{id:int} - specific types
    /people/{id:string}
*/
type GowebPath struct {
	RawPath      string
	regexPattern string
	path         *Path
}

func NewGowebPath(path string) (*GowebPath, error) {

	p := new(GowebPath)
	p.RawPath = path
	p.path = NewPath(path)

	return p, nil
}

func (p *GowebPath) GetPathMatch(path *Path) *PathMatch {

	pathMatch := new(PathMatch)
	pathMatch.Matches = true

	checkSegments := p.path.Segments()
	pathSegments := path.Segments()

	// make sure the segments match in length, or there is a catchall
	// at the end of the check path
	if len(checkSegments) < len(pathSegments) {

		// is the last segment a catchall?
		if getSegmentType(checkSegments[len(checkSegments)-1]) != segmentTypeCatchall {
			return PathDoesntMatch
		}

	} else if len(checkSegments) > len(pathSegments) {

		if getSegmentType(checkSegments[len(checkSegments)-1]) != segmentTypeDynamicOptional {
			return PathDoesntMatch
		}

	}

	// check each segment
	for segmentIndex, checkSegment := range checkSegments {

		switch getSegmentType(checkSegment) {
		case segmentTypeLiteral:

			pathSegment := pathSegments[segmentIndex]
			if strings.ToLower(pathSegment) != strings.ToLower(checkSegment) {
				return PathDoesntMatch
			}

		case segmentTypeDynamicOptional:

		}

	}

	return pathMatch

}

func (p *GowebPath) RegexPattern() string {
	return p.regexPattern
}