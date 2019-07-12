package models

type SplitPayment struct {
  Email string `json:"email"`
  FixedCut int `json:"fixed_cut"`
  VariableCut string `json:"variable_cut"`
  StackOrder int `json:"stack_order"`
}

type Collection struct {
  Title string `json:"title"`
  SplitHeader bool `json:"split_header"`
  SplitPayments []SplitPayment `json:"split_payments"`
}
