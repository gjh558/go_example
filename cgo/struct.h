#ifndef LIBADS_PACING_H
#define LIBADS_PACING_H

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdint.h>
#include <stdlib.h>

struct YieldOpt
{
  int32_t len_;
  int32_t* ids_;
};

struct Counter {
  int64_t id_;
  int64_t imp_;
  int64_t cur_;
  int64_t targeted_imp_;
  int64_t targeted_cur_;
  int64_t daily_imp_;
  int64_t daily_cur_;
  int64_t price_;
};

struct Candidate {
  int64_t id_; // should be ads entity id
  int32_t magnifier_;
  float iso_threshold_;
  struct YieldOpt yield_opt_;
  // struct Counter* counter_;
};

struct Context{
  int32_t flags_;
  float one2n_pacing_threshold_;
  int64_t timestamp_;
  int64_t cro_;
  int64_t co_;
  int64_t dist_;
};

struct Record {
  int32_t status_;
  float life_stage_;
  float gross_fill_rate_;
  float ext_gross_fill_rate_;
  float targted_fill_rate_;
  float ext_on_targeted_fill_rate_;
  float daily_fill_rate_;
};

struct PacingRequest {
  struct Context context_;
  int32_t cands_len_;
  int32_t counter_len_;
  struct Candidate* cands_;
  struct Counter* counters_;
};

struct PacingResponse {
  struct Record* records_;
};

struct Pacing {
  struct PacingRequest *request_;
  struct PacingResponse *response_;
};

uint32_t* alloc() {
  uint32_t* p = (uint32_t*)malloc(sizeof(uint32_t) * 10);
  p[0] = 9;
  for (int i = 1; i < 10; ++i)
    p[i] = i;
  return p;
}

// struct PacingRequest* new_pacing_request(int cand_size, int counter_size)
// {
//   int counter_bytes = sizeof(struct Counter)*counter_size;
//   int cand_bytes = sizeof(struct Candidate)*cand_size;
//   int total_bytes = sizeof(struct PacingRequest) + cand_bytes + counter_bytes;
//   struct PacingRequest* req = malloc(total_bytes);
//   req->cands_ = (struct Candidate*)(req + 1);
//   req->counters_ = (struct Counter*)(req->cands_ + 1);
//   return req;
// }
//
// struct PacingResponse* new_pacing_response(int cand_size)
// {
//   return malloc(sizeof(struct PacingResponse) + sizeof(struct Record)*cand_size);
// }
//
// struct Pacing* new_pacing(int cand_size, int counter_size)
// {
//   struct Pacing* p = malloc(sizeof(struct Pacing));
//   p->request_ = new_pacing_request(cand_size, counter_size);
//   p->request_->cands_->id_ = 10;
//   p->request_->cands_len_ = 1;
//   p->response_ = new_pacing_response(cand_size);
//   return p;
// }
// void free_pacing(struct Pacing *p)
// {
//   free(p->response_);
//   free(p->request_);
//   free(p);
// }

#ifdef __cplusplus
} //extern "C"
#endif

#endif // LIBADS_PACING_H
