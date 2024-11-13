package disjointsets


type disjointset struct{
	components []int
	component_size []int
}

func NewDisjointSet(n int) disjointset {
	comps := make([]int, n);
	comps_size := make([]int, n);
	for i := 0; i < n; i++ {
		comps[i] = i;
		comps_size[i] = 1;	
	}
	ds := disjointset{
		components: comps,
		component_size: comps_size,
	}
	return ds;
}

func getObject(ds disjointset, i int) int{
	return ds.components[i];
}

func GetRoot(ds disjointset, i int) int {
	i_root := getObject(ds,i);
	for i_root != i{
		ds.components[i] = getObject(ds,i_root);
		i = i_root;
		i_root = getObject(ds,i);
	} 
	return i_root;
}

func IsConnected(ds disjointset, p, q int) bool {
	return GetRoot(ds,p) == GetRoot(ds,q);
}

func Union(ds disjointset, p, q int) {
	p_root := GetRoot(ds, p);
	q_root := GetRoot(ds, q);
	if ds.component_size[p_root] > ds.component_size[q_root]{
		ds.components[p_root] = q_root;
		ds.component_size[p_root] += ds.component_size[q_root];
	}else{
		ds.components[q_root] = p_root;
		ds.component_size[q_root] += ds.component_size[p_root]; 
	}
}