package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Items       []TaskItem `json:"items" gorm:"foreignKey:Task_ID"` // Establishes the relationship with TaskItem
}

func (Task) TableName() string {
	return "task" // Specify the custom table name here
}

type TaskItem struct {
	ID      int    `json:"id"`
	Task_ID int    `json:"task_id"`
	Item    string `json:"description"`
}

func (TaskItem) TableName() string {
	return "task_item" // Specify the custom table name here
}

type TaskDto struct {
	DB *gorm.DB
}

func (taskDto TaskDto) GetAllTasks() ([]Task, error) {
	var task []Task

	// Using GORM to perform the equivalent query with eager loading of associated items
	result := taskDto.DB.Preload("Items").Find(&task).Unscoped()

	if result.Error != nil {
		return nil, result.Error
	}

	// The Find() method has executed the query and loaded associated items

	return task, nil
}

/* func (taskDto TaskDto) GetAllTasks() ([]Task, error) {

	query := `
		SELECT t.id, t.title, t.description, t.completed, t.created_at, t.updated_at, ti.item
		FROM task t
		LEFT JOIN task_item ti ON t.id = ti.task_id
		ORDER BY t.id, ti.id
	`

	rows, err := taskDto.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskMap := make(map[int]*Task)

	taskItemsMap := make(map[int][]string)

	for rows.Next() {
		var taskItem sql.NullString
		task := &Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
			&taskItem,
		)
		if err != nil {
			return nil, err
		}

		if taskItem.Valid {
			taskItemsMap[task.ID] = append(taskItemsMap[task.ID], taskItem.String)
		}

		if _, ok := taskMap[task.ID]; !ok {
			taskMap[task.ID] = task
		}
	}

	for taskID, taskItems := range taskItemsMap {
		if task, ok := taskMap[taskID]; ok {
			task.Items = taskItems
		}
	}

	tasks := make([]Task, 0, len(taskMap))
	for _, task := range taskMap {
		tasks = append(tasks, *task)
	}

	return tasks, nil
} */

func (taskDto TaskDto) Insert(task *Task) error {
	// Using GORM to perform the equivalent insert operation
	result := taskDto.DB.Create(task).Unscoped()
	if result.Error != nil {
		return result.Error
	}

	return nil
}

/* func (taskDto TaskDto) Insert(task *Task) error {

	stmt, err := taskDto.DB.Prepare(`
			INSERT INTO task (title, description, completed, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var taskID int
	err = stmt.QueryRow(task.Title, task.Description, task.Completed, task.CreatedAt, task.UpdatedAt).Scan(&taskID)
	if err != nil {
		return err
	}

	task.ID = taskID

	return nil

} */

func (taskDto TaskDto) Get(id int64) (*Task, error) {
	return nil, nil
}

func (taskDto TaskDto) UpdateTask(id int, task *Task) error {
	// Using GORM to perform the equivalent update operation
	result := taskDto.DB.Model(&Task{}).Where("id = ?", id).Updates(Task{
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		UpdatedAt:   time.Now(),
		Items:       task.Items, // Update the Items field directly
	}).Unscoped()
	if result.Error != nil {
		return result.Error
	}

	// Delete existing task items
	err := taskDto.DB.Where("task_id = ?", id).Delete(&TaskItem{}).Unscoped().Error
	if err != nil {
		return err
	}

	// Insert new task items
	for _, item := range task.Items {
		taskItem := TaskItem{
			Task_ID: id,
			Item:    item.Item,
		}
		err := taskDto.DB.Create(&taskItem).Error
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	 func (taskDto TaskDto) UpdateTask(id int, task *Task) error {

		stmt, err := taskDto.DB.Prepare(`
			UPDATE task
			SET title = $1, description = $2, completed = $3, updated_at = $4
			WHERE id = $5
		`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(task.Title, task.Description, task.Completed, time.Now(), id)
		if err != nil {
			return err
		}

		_, err = taskDto.DB.Exec("DELETE FROM task_item WHERE task_id = $1", task.ID)
		if err != nil {
			return err
		}

		stmt, err = taskDto.DB.Prepare("INSERT INTO task_item (task_id, item) VALUES ($1, $2)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, item := range task.Items {
			_, err := stmt.Exec(task.ID, item)
			if err != nil {
				return err
			}
		}
		return nil
	}
*/

func (taskDto TaskDto) DeleteTask(id int) error {
	// Using GORM to perform the delete operation
	result := taskDto.DB.Where("id = ?", id).Delete(&Task{}).Unscoped()
	if result.Error != nil {
		return result.Error
	}

	/* // You might want to also delete associated items for the task
	err := taskDto.DB.Where("task_id = ?", id).Delete(&TaskItem{}).Error
	if err != nil {
		return err
	} */

	return nil
}

/* func (taskDto TaskDto) DeleteTask(id int) error {

	stmt, err := taskDto.DB.Prepare(`
		DELETE FROM task WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
} */

func (taskDto TaskDto) InsertTaskItem(taskID int, item string) error {
	taskItem := TaskItem{
		Task_ID: taskID,
		Item:    item,
	}

	err := taskDto.DB.Create(&taskItem).Error
	if err != nil {
		return err
	}

	return nil
}

/* func (taskDto TaskDto) InsertTaskItem(taskID int, item string) error {

	stmt, err := taskDto.DB.Prepare(`
		INSERT INTO task_item (task_id, item)
		VALUES ($1, $2)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskID, item)
	if err != nil {
		return err
	}

	return nil
} */

/* func (taskDto TaskDto) GetTask(id int) (*Task, error) {
	query := `
		SELECT t.id, t.title, t.description, t.completed, t.created_at, t.updated_at, ti.item
		FROM task t
		LEFT JOIN task_item ti ON t.id = ti.task_id
		WHERE t.id = $1
	`

	rows, err := taskDto.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	task := &Task{}
	task.Items = make([]string, 0)

	for rows.Next() {
		var taskItem sql.NullString
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,
			&taskItem,
		)
		if err != nil {
			return nil, err
		}

		if taskItem.Valid {
			task.Items = append(task.Items, taskItem.String)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return task, nil
} */

func (taskDto TaskDto) GetTask(id int) (*Task, error) {
	task := &Task{}

	err := taskDto.DB.Preload("Items").First(task, id).Unscoped().Error
	if err != nil {
		return nil, err
	}

	return task, nil
}
