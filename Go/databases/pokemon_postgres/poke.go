package pokemon_postgres

import (
	t "Go/time_completion"
	"database/sql"
	"strconv"
)

<<<<<<< Updated upstream
type Pokemon struct {
	//gorm.Model
	dexnum int
	name string
	type1 string
	type2 string
	stage string
	evolve_level int
	gender_ratio string
	height float32
	weight float32
	description string
	category string
	lvl_speed float32
	base_exp int
	catch_rate int
=======
<<<<<<< Updated upstream


// var ctx context.Context
// var cli client.Client

func main() {
	err := startColima()
=======
type Pokemon struct {
	//gorm.Model
	Dexnum       int
	Name         string
	Type1        string
	Type2        string
	Stage        string
	Evolve_level int
	Gender_ratio string
	Height       float32
	Weight       float32
	Description  string
	Category     string
	Lvl_speed    float32
	Base_exp     int
	Catch_rate   int
}

//func (Pokemon) TableName() string {
//	return "Pokedex"
//}

//func GormAllPokemon(db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{}
//	allPokemon := db.Find(&poke)
//	if allPokemon.Error != nil {
//		return "", allPokemon.Error
//	}
//	println(allPokemon.RowsAffected)
//
//	return fmt.Sprintf("%+v\n", allPokemon), nil
//}

//func GormAPokemon(Dexnum int, db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{Dexnum: Dexnum}
//
//	resultPoke := db.First(&poke)
//	if resultPoke.Error != nil {
//		return "", resultPoke.Error
//	}
//
//
//	return fmt.Sprintf("%+v\n", poke), nil
//}

func GetAllPokemon(db *sql.DB) ([]Pokemon, error) {
	defer t.Timer()()
	rows, err := db.Query(`SELECT * FROM "Pokedex";`)
	defer rows.Close()
>>>>>>> Stashed changes
	if err != nil {
		fmt.Printf("Colima Error: %s\n", erp.Cause(err).Error())
		os.Exit(1)
	}
	
	ctx, cli, err := connect()
	if err != nil{
		fmt.Printf("Fatal Startup Error -> Cause: %s\n", erp.Cause(err).Error())
		os.Exit(1)
	}
    defer cli.Close()

<<<<<<< Updated upstream
	// pull(cli, ctx, "postgres")
	// listImages(cli, ctx)
	// i, err := getImage(cli,ctx, "postgres")
	// if err != nil{
	// 	fmt.Printf("%+v\n", erp.Cause(err))
	// }
	// fmt.Printf("Image: %+v\n", i)


	err = postgresContainer(cli, ctx)
	if err != nil{
		println("err: ", erp.Cause(err).Error())
	}

	println("Waiting 20s")
	time.Sleep(time.Duration(time.Second) * 20)
	postgres, err := getContainer(cli, ctx, "postgres")
	if err != nil{
		println("err: ", erp.Cause(err).Error())
	}
	err = deleteContainer(cli, ctx, postgres.ID)
	if err != nil {
		println("err: ", erp.Cause(err).Error())
	}



	// pruneContainer(cli, ctx)
	// listContainers(cli,ctx, true)
	// err = stop(cli, ctx, id)
	// if err != nil{
	// 	println("err: ", erp.Cause(err))
	// }

	// err = deleteContainer(cli, ctx, id)
	// if err != nil{
	// 	println("err: ", erp.Cause(err))
	// }

    // out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    // if err != nil {
    //     panic(err)
    // }

    // stdcopy.StdCopy(os.Stdout, os.Stderr, out)
>>>>>>> Stashed changes
}
//func (Pokemon) TableName() string {
//	return "Pokedex"
//}

//func GormAllPokemon(db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{}
//	allPokemon := db.Find(&poke)
//	if allPokemon.Error != nil {
//		return "", allPokemon.Error
//	}
//	println(allPokemon.RowsAffected)
//
//	return fmt.Sprintf("%+v\n", allPokemon), nil
//}

//func GormAPokemon(dexnum int, db *gorm.DB)(string, error){
//	defer t.Timer()()
//	var poke = Pokemon{dexnum: dexnum}
//
//	resultPoke := db.First(&poke)
//	if resultPoke.Error != nil {
//		return "", resultPoke.Error
//	}
//
//
//	return fmt.Sprintf("%+v\n", poke), nil
//}

func GetAllPokemon(db *sql.DB) ([]Pokemon, error) {
	defer t.Timer()()
	rows, err := db.Query(`SELECT * FROM "Pokedex";`)
	defer rows.Close()
	if err != nil {
<<<<<<< Updated upstream
		println("Query Fail")
		return nil, err
=======
		//println(erp.Cause(err).Error())
		cmd := exec.Command("colima", "start")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
=======
	for rows.Next() {
		//err := rows.Scan(&poke) splat operator?
		err := rows.Scan(&poke.Dexnum, &poke.Name, &poke.Type1, &poke.Type2, &poke.Stage, &poke.Evolve_level, &poke.Gender_ratio, &poke.Height, &poke.Weight, &poke.Description, &poke.Category, &poke.Lvl_speed, &poke.Base_exp, &poke.Catch_rate)
>>>>>>> Stashed changes
		if err != nil {
			return erp.Wrap(err, "Failed to start Colima\n")
		}
	}
	return nil
}

func getUser() (string, error){
	files, err := ioutil.ReadDir("/Users/")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
		if s.HasPrefix(f.Name(), "Z") && len(f.Name()) == 7 {
			return f.Name(), nil
		}
    }
	return "", erp.New("cannot find User zID folder in system")
}

func listContainers(cli *client.Client, ctx context.Context, all bool) error{
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: all})
	if err != nil {
		return erp.Wrap(err, "failed to get containerList")
	}

	fmt.Printf("%-12s\t%-30s\t%-10s\t%s\n", "CONTAINER ID", "IMAGE", "STATUS", "PORTS")
	for _, c := range containers {
		fmt.Printf("%-12s\t%.30s\t%.10s\t%v\n", c.ID[:12], fmt.Sprintf("%-30s",c.Image), c.Status, c.Ports )
	}
	return nil
}

func getContainer(cli *client.Client, ctx context.Context, nameOrId string) (types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	ct := types.Container{}
	if err != nil {
		return ct, erp.Wrap(err, "failed to get containerList")
>>>>>>> Stashed changes
	}
	var poke Pokemon
	var allpoke []Pokemon

	for rows.Next() {
		//err := rows.Scan(&poke) splat operator?
		err := rows.Scan(&poke.dexnum, &poke.name, &poke.type1, &poke.type2, &poke.stage, &poke.evolve_level, &poke.gender_ratio, &poke.height, &poke.weight, &poke.description, &poke.category, &poke.lvl_speed, &poke.base_exp, &poke.catch_rate)
		if err != nil {
			println("Scan fail")
			return nil, err
		}
		allpoke = append(allpoke, poke)
	}
	return allpoke, nil
}

<<<<<<< Updated upstream
func GetAPokemon(dexnum string, db *sql.DB) (*Pokemon, error) {
	defer t.Timer()()
	atoi, err := strconv.Atoi(dexnum)
	if err != nil || atoi < 1{
		println("Invalid Dexnum")
		return nil, err
=======
<<<<<<< Updated upstream
func stop(cli *client.Client, ctx context.Context, id string) error{
	err := cli.ContainerStop(ctx, id, nil)
	if err != nil{
		return erp.Wrap(err, "Failed to stop container")
>>>>>>> Stashed changes
	}
	var poke Pokemon
	query := `SELECT * FROM "Pokedex" WHERE dexnum=` + dexnum
	rows := db.QueryRow(query)

<<<<<<< Updated upstream
	err = rows.Scan(&poke.dexnum, &poke.name, &poke.type1, &poke.type2, &poke.stage, &poke.evolve_level, &poke.gender_ratio, &poke.height, &poke.weight, &poke.description, &poke.category, &poke.lvl_speed, &poke.base_exp, &poke.catch_rate)
=======
func start(cli *client.Client, ctx context.Context, id string) error{
	opts := types.ContainerStartOptions{

	}
	err := cli.ContainerStart(ctx, id, opts)
	if err != nil{
		return erp.Wrap(err, "Failed to start container")
	}

	// statusCh, errCh := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)
    // select {
    // case err := <-errCh:
    //     if err != nil {
    //         panic(err)
    //     }
    // case <-statusCh:
		
    // }
	// return nil
	return nil
}

func listImages(cli *client.Client, ctx context.Context) error{
	opts := types.ImageListOptions{
		All: true,	
	}
	images, err := cli.ImageList(ctx, opts)
=======
func GetAPokemon(dexnum string, db *sql.DB) (*Pokemon, error) {
	defer t.Timer()()
	atoi, err := strconv.Atoi(dexnum)
	if err != nil || atoi < 1 {
		println("Invalid Dexnum")
		return nil, err
	}
	var poke Pokemon
	query := `SELECT * FROM "Pokedex" WHERE Dexnum=` + dexnum
	rows := db.QueryRow(query)

	err = rows.Scan(&poke.Dexnum, &poke.Name, &poke.Type1, &poke.Type2, &poke.Stage, &poke.Evolve_level, &poke.Gender_ratio, &poke.Height, &poke.Weight, &poke.Description, &poke.Category, &poke.Lvl_speed, &poke.Base_exp, &poke.Catch_rate)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	if err != nil {
		println("Scan fail")
		return nil, err
	}
<<<<<<< Updated upstream
	return &poke, nil
}
=======
<<<<<<< Updated upstream
	return nil
}

func getImage(cli *client.Client, ctx context.Context, nameOrId string) (types.ImageSummary,error){
	opts := types.ImageListOptions{
		All: true,	
	}
	images, err := cli.ImageList(ctx, opts)

	t := types.ImageSummary{}

	if err != nil {
		return t, err
	}
	for _, i := range images {
		if i.ID == nameOrId {
			return i,nil
		}
		for _, name := range i.RepoTags{
			if s.Contains(name, nameOrId){
				return i, nil
			}
		}
	}
	return t, nil
}

func create(cli *client.Client, ctx context.Context, tainerConfig *container.Config, hostCon *container.HostConfig,
	netConfig *network.NetworkingConfig, platform *v1.Platform, cName string) (string, error){
    cont, err := cli.ContainerCreate(ctx, tainerConfig, hostCon, netConfig, platform, cName)
    if err != nil {
        return "", erp.Wrap(err, "Failed to create Container")
    }
	return cont.ID, nil
}

//func postgresContainer(cli *client.Client, ctx context.Context) (tainerConfig *container.Config, hostCon *container.HostConfig,
//	netConfig *network.NetworkingConfig, platform *v1.Platform, cName string, err error){
	// func postgresContainer(cli *client.Client, ctx context.Context) (*container.Config, *container.HostConfig,
		// *network.NetworkingConfig, *v1.Platform, string, error){
func postgresContainer(cli *client.Client, ctx context.Context) error{
	// expPort := "5000"
	cName := "postgres"
	// container = client.containers.run(
    //     image_name,
    //     detach=True,
    //     name=container_name,
    //     ports={'5000/tcp': 5000}
    //     # environment={"POSTGRES_PASSWORD": "pokemon", "POSTGRES_DB": "Pokemon"}
    // )
	
	// newport, err := nat.NewPort("tcp", expPort)
	// if err != nil{
	// 	erp.Wrap(err, fmt.Sprintf("Unable to create docker port: %s", newport.Port()))
	// }

	// exposedPorts := map[nat.Port]struct{}{
	// 	newport: struct{}{},
	// }

	tainerConfig := container.Config{
		Image: "postgres",
		Env: []string{"POSTGRES_PASSWORD=password"},
		// ExposedPorts: exposedPorts,//List of exposed ports

		// AttachStdin: true, //Attach the standard input, makes possible user interaction

		// Cmd:   []string{"echo", "hello world"},
		// WorkingDir: "/",
		}

	// *hostCon = container.HostConfig{
	// 	PortBindings: nat.PortMap{
	// 		newport: []nat.PortBinding{
	// 			{
	// 				HostIP:   "0.0.0.0",
	// 				HostPort: expPort,
	// 			},
	// 		},
	// 	},
	// 	RestartPolicy: container.RestartPolicy{
	// 		Name: "always",
	// 	},
	// 	LogConfig: container.LogConfig{
	// 		Type:   "json-file",
	// 		Config: map[string]string{},
	// 	},
	// }

	// *netConfig = network.NetworkingConfig{
	// 	EndpointsConfig: map[string]*network.EndpointSettings{},
	// }

	// platform = &v1.Platform{}

	// return &tainerConfig, nil, nil, nil, cName, nil

	fmt.Printf("Creating Container %s\n", cName)
	id, err := create(cli, ctx, &tainerConfig, nil, nil, nil, cName)
	if err != nil{
		return err
	}
	fmt.Printf("Starting Container %s, id: %s\n", cName, id[:5])
	err = start(cli, ctx, id)
	if err != nil{
		return err
	}
	return nil
=======
	return &poke, nil
>>>>>>> Stashed changes
}
>>>>>>> Stashed changes
