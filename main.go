//processService contiene las reglas del servicio
func(s store) processService(data *model) ([][]int, error) {
	idx := len(data.Matriz)
	arr := make([][]int, idx)

	//condicion NxN && iniciar la matriz secundaria
	for i := 0; i < idx; i++ {
		if len(data.Matriz[i]) != idx {
			return nil, errors.New("Matriz no valida, solo se aceptan matrices NxN")
		}
		arr[i] = make([]int, idx)
	}

	//rotar matriz en 90°
	for i := 0; i < idx; i++ {
		for j := 0; j < idx; j++ {
			arr[j][i] = data.Matriz[i][idx-j-1]
		}
	}

	return arr, nil
}