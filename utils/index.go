package utils

type Index map [string][]int

func (idx Index) Add(docs []document){
	for _, doc := range(docs){
		for _, token := range analyze(doc.Text){
			ids := idx[token]
			if ids != nil && ids[len(ids) - 1] == doc.ID{
				continue
			}
			idx[token] = append(idx[token], doc.ID)
		}
	}
}