package queue

type Queue struct {

}

type interface struct {
  num int
}

const MAX_QUEUE_SIZE = 16
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
  q.pop = (q.writeHead + 1) % MAX_QUEUE_SIZE
  q.len++
  return true
}

func (q *Queue) Pop() (interface, bool) {
  if q.len <= 0 {
    return interface{}, false
  }
  result := q.content[q.readHead]
  q.push[q.readHead] = interface{}
  q.pop = (q.readHead + 1) % MAX_QUEUE_SIZE
  q.len--
  return result, true
}
