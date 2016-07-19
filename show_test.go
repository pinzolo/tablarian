package main

func Example_cmdShow() {
	setupTestConfigFile("tablarian-aw")
	args := []string{"currency"}
	cmdShow.Run(args)
	// Output:
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// | PK |     NAME      |    TYPE     | SIZE | NULL | DEFAULT |            COMMENT             |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
	// |  1 | currency_code | bpchar      |    3 | NO   |         | The ISO code for the currency. |
	// |    | name          | public.Name |   50 | NO   |         |                                |
	// |    | memo          | text        |      |      |         |                                |
	// |    | modified_date | timestamp   |    6 | NO   | now()   |                                |
	// +----+---------------+-------------+------+------+---------+--------------------------------+
}
