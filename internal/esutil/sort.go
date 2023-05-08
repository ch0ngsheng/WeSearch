package esutil

func (r SearchDocResp) Len() int {
	return len(r.DocIDList)
}

func (r SearchDocResp) Less(i, j int) bool {
	return r.DocIDList[i].Score > r.DocIDList[j].Score
}

func (r SearchDocResp) Swap(i, j int) {
	r.DocIDList[i], r.DocIDList[j] = r.DocIDList[j], r.DocIDList[i]
}
