package queue

type interface struct {
  num int
}

const MAX_QUEUE_SIZE = 16 //fixed size

type Queue struct {
  push key interface{}
  pop interface{}
  contains interface{} bool
  len int
  keys [MAX_QUEUE_SIZE]interface{}
}

func (q *Queue) Push(e interface) bool {
  if q.len >= MAX_QUEUE_SIZE {
    return false
  }
  q.push[q.keys] = e
  q.pop = (q.contains + 1) % MAX_QUEUE_SIZE
  q.len++
  return true
}
