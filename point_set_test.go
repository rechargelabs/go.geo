package geo

import (
	"math"
	"testing"
)

func TestGepCentroid(t *testing.T) {
	ps := &PointSet{}
	ps.Push(&Point{-122.42558918, 37.76159786}).
		Push(&Point{-122.41486043, 37.78138826}).
		Push(&Point{-122.40206146, 37.77962363})
	centroid := ps.GeoCentroid()
	expectedCenter := &Point{-122.41417035666666, 37.77420325}
	if !centroid.Equals(expectedCenter) {
		t.Errorf("should find centroid correctly, got %v", centroid.Lng())
	}
}

func TestGeoDistanceFrom(t *testing.T) {
	ps := &PointSet{}
	ps.Push(&Point{-122.42558918, 37.76159786}).
		Push(&Point{-122.41486043, 37.78138826}).
		Push(&Point{-122.40206146, 37.77962363})
	fromPoint := &Point{-122.41941550000001, 37.7749295}

	if distance := ps.GeoDistanceFrom(fromPoint); math.Floor(distance) != 823 {
		t.Errorf("should find geo distance from correctly, got %v", distance)
	}
}

func TestNewPointSet(t *testing.T) {
	ps := NewPointSet()
	ps.Push(&Point{-122.42558918, 37.76159786}).
		Push(&Point{-122.41486043, 37.78138826}).
		Push(&Point{-122.40206146, 37.77962363})
	if ps.Length() != 3 {
		t.Errorf("should find correct length of new point set %v", ps.Length())
	}
}

func TestNewPointSetPreallocate(t *testing.T) {
	ps := NewPointSet()
	ps.Push(&Point{-122.42558918, 37.76159786}).
		Push(&Point{-122.41486043, 37.78138826}).
		Push(&Point{-122.40206146, 37.77962363})

	if ps.Length() != 3 {
		t.Errorf("should find correct length of new point set %v", ps.Length())
	}

	if !ps.GetAt(0).Equals(&Point{-122.42558918, 37.76159786}) {
		t.Errorf("should find correct first point of new point set %v", ps.GetAt(0))
	}

	if !ps.GetAt(2).Equals(&Point{-122.40206146, 37.77962363}) {
		t.Errorf("should find correct first point of new point set %v", ps.GetAt(2))
	}
}

func TestPathBound(t *testing.T) {
	ps := NewPointSet()
	ps.Push(NewPoint(0.5, .2))
	ps.Push(NewPoint(-1, 0))
	ps.Push(NewPoint(1, 10))
	ps.Push(NewPoint(1, 8))

	answer := NewBound(-1, 1, 0, 10)
	if b := ps.Bound(); !b.Equals(answer) {
		t.Errorf("bound, %v != %v", b, answer)
	}

	ps = NewPointSet()
	if !ps.Bound().Empty() {
		t.Error("expect empty point set to have empty bounds")
	}
}

func TestPointSetSetAt(t *testing.T) {
	ps := NewPointSet()
	point := NewPoint(1, 2)

	ps.Push(NewPoint(2, 3))

	ps.SetAt(0, point)
	if p := ps.GetAt(0); !p.Equals(point) {
		t.Errorf("setAt expected %v == %v", p, point)
	}
}

func TestPointSetSetAtPanicIndexOver(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect setAt to panic if index out of range")
		}
	}()

	ps := NewPointSet()
	ps.Push(NewPoint(1, 2))
	ps.SetAt(2, NewPoint(3, 4))
}

func TestPointSetSetAtPanicIndexUnder(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect setAt to panic if index out of range")
		}
	}()

	ps := NewPointSet()
	ps.Push(NewPoint(1, 2))
	ps.SetAt(-1, NewPoint(3, 4))
}

func TestPointSetGetAt(t *testing.T) {
	ps := NewPointSet()
	point := NewPoint(1, 2)

	ps.Push(point)

	if p := ps.GetAt(0); !p.Equals(point) {
		t.Errorf("getAt expected %v == %v", p, point)
	}

	if p := ps.GetAt(10); p != nil {
		t.Error("expect out of range getAt to be nil")
	}

	if p := ps.GetAt(-1); p != nil {
		t.Error("expect negative index getAt to be nil")
	}

	if p := ps.GetAt(0).SetX(100); !p.Equals(ps.GetAt(0)) {
		t.Error("expect getAt to return pointer to original value")
	}
}

func TestPointSetInsertAt(t *testing.T) {
	ps := NewPointSet()
	point1 := NewPoint(1, 2)
	point2 := NewPoint(3, 4)
	ps.Push(point1)

	ps.InsertAt(0, point2)
	if p := ps.GetAt(0); !p.Equals(point2) {
		t.Errorf("insertAt expected %v == %v", p, point2)
	}

	if p := ps.GetAt(1); !p.Equals(point1) {
		t.Errorf("insertAt expected %v == %v", p, point1)
	}

	point3 := NewPoint(5, 6)
	ps.InsertAt(2, point3)
	if p := ps.GetAt(2); !p.Equals(point3) {
		t.Errorf("insertAt expected %v == %v", p, point3)
	}
}

func TestPointSetInsertAtPanicIndexOver(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect insertAt to panic if index out of range")
		}
	}()

	ps := NewPointSet()
	ps.Push(NewPoint(1, 2))
	ps.InsertAt(2, NewPoint(3, 4))
}

func TestPointSetInsertAtPanicIndexUnder(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect insertAt to panic if index out of range")
		}
	}()

	ps := NewPointSet()
	ps.Push(NewPoint(1, 2))
	ps.InsertAt(-1, NewPoint(3, 4))
}

func TestPointSetRemoveAt(t *testing.T) {
	ps := NewPointSet()
	point := NewPoint(1, 2)

	ps.Push(point)
	ps.RemoveAt(0)

	if ps.Length() != 0 {
		t.Error("expect removeAt to remove point")
	}
}

func TestPointSetRemoveAtPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expect removeAt to panic if index out of range")
		}
	}()

	ps := NewPointSet()
	ps.Push(NewPoint(1, 2))
	ps.RemoveAt(2)
}

func TestPointSetPush(t *testing.T) {
	ps := NewPointSet()
	ps.Push(NewPoint(1, 2))

	if ps.Length() != 1 {
		t.Errorf("push length 1 != %d", ps.Length())
	}

	answer := NewPoint(1, 2)
	if a := ps.GetAt(0); !a.Equals(answer) {
		t.Errorf("push first expecting %v == %v", a, answer)
	}
}

func TestPointSetPop(t *testing.T) {
	ps := NewPointSet()

	if ps.Pop() != nil {
		t.Error("expect empty pop to return nil")
	}

	ps.Push(NewPoint(1, 2))
	answer := NewPoint(1, 2)
	if a := ps.Pop(); !a.Equals(answer) {
		t.Errorf("pop first expecting %v == %v", a, answer)
	}
}

func TestPointSetEquals(t *testing.T) {
	p1 := NewPointSet()
	p1.Push(NewPoint(0.5, .2))
	p1.Push(NewPoint(-1, 0))
	p1.Push(NewPoint(1, 10))

	p2 := NewPointSet()
	p2.Push(NewPoint(0.5, .2))
	p2.Push(NewPoint(-1, 0))
	p2.Push(NewPoint(1, 10))

	if !p1.Equals(p2) {
		t.Error("equals paths should be equal")
	}

	p3 := p2.Clone().SetAt(1, NewPoint(0, 0))
	if p1.Equals(p3) {
		t.Error("equals paths should not be equal")
	}

	p2.Pop()
	if p2.Equals(p1) {
		t.Error("equals paths should not be equal")
	}
}

func TestPathClone(t *testing.T) {
	p1 := NewPointSet()
	p1.Push(NewPoint(0, 0))
	p1.Push(NewPoint(0.5, .2))
	p1.Push(NewPoint(1, 0))

	p2 := p1.Clone()
	p2.Pop()
	if p1.Length() == p2.Length() {
		t.Errorf("clone length %d == %d", p1.Length(), p2.Length())
	}

	p2 = p1.Clone()
	if p1 == p2 {
		t.Error("clone should return different pointers")
	}

	if !p2.Equals(p1) {
		t.Error("clone paths should be equal")
	}
}