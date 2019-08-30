package scheduler

import "gopkg.in/robfig/cron.v3"

type SchedulerStore struct {
  cron    *cron.Cron
  Running  bool
}

func BuildStore() SchedulerStore {

  return SchedulerStore {
    cron: cron.New(),
  }
}

func (s *SchedulerStore) Start() {

  if !s.Running {
    s.cron.Start()
    s.Running = true
  }
}

func (s *SchedulerStore) Stop() {

  if s.Running {
    s.cron.Stop()
    s.Running = false
  }
}

func (s *SchedulerStore) AddJob(cron string, job cron.Job) {

  s.cron.AddJob(cron, job)
}

func (s *SchedulerStore) AddFunction(cron string, fn cron.FuncJob) {

  s.cron.AddFunc(cron, fn)
}

func (s *SchedulerStore) Entries() []cron.Entry {

    return s.cron.Entries()
}

func (s *SchedulerStore) GetEntry(id int) cron.Entry {

  return s.cron.Entry(cron.EntryID(id))
}

func (s *SchedulerStore) RemoveEntry(id int) {

  s.cron.Remove(cron.EntryID(id))
}

func (s *SchedulerStore) RemoveAllEntries() {

  for _, entry := range s.Entries() {

    s.RemoveEntry(int(entry.ID))
  }
}
