package board

// Represent the pegs using mere strings. One step closer
// to representing it with binary representation.
// The following behavior is used to accomplish this, which
// is even on its own a preliminary attempt at doing a
// string representation of the map.
// The strings do require environment/global variables to be
// set to parse them, and equal strings can represent different
// boards under different environment variables
//
// a Board like this:
//   0 1 1 1 1 1 0
//   0 0 0 1 0 0 0
//   0 0 0 1 0 0 0
//   0 0 0 1 0 0 0
// 0 0 0 0 0 0 0
//
