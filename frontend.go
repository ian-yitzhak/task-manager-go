package main

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go CRUD - Task Manager</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
            background: #1a1a2e;
            color: #eee;
            min-height: 100vh;
            display: flex;
            justify-content: center;
            padding: 40px 20px;
        }
        .container { width: 100%; max-width: 600px; }
        h1 {
            text-align: center;
            margin-bottom: 30px;
            color: #e94560;
            font-size: 2rem;
        }
        .add-form {
            display: flex;
            gap: 10px;
            margin-bottom: 30px;
        }
        .add-form input {
            flex: 1;
            padding: 12px 16px;
            border: 2px solid #16213e;
            border-radius: 8px;
            background: #16213e;
            color: #eee;
            font-size: 1rem;
            outline: none;
            transition: border-color 0.2s;
        }
        .add-form input:focus { border-color: #e94560; }
        .add-form button {
            padding: 12px 24px;
            background: #e94560;
            color: white;
            border: none;
            border-radius: 8px;
            font-size: 1rem;
            cursor: pointer;
            transition: background 0.2s;
        }
        .add-form button:hover { background: #c23152; }
        .task-list { list-style: none; }
        .task-item {
            display: flex;
            align-items: center;
            gap: 12px;
            padding: 14px 16px;
            background: #16213e;
            border-radius: 8px;
            margin-bottom: 10px;
            transition: transform 0.1s;
        }
        .task-item:hover { transform: translateX(4px); }
        .task-item.done .task-title {
            text-decoration: line-through;
            opacity: 0.5;
        }
        .task-title { flex: 1; font-size: 1rem; }
        .task-item input[type="checkbox"] {
            width: 20px;
            height: 20px;
            cursor: pointer;
            accent-color: #e94560;
        }
        .btn-delete {
            padding: 6px 14px;
            background: #533a3a;
            color: #e94560;
            border: none;
            border-radius: 6px;
            cursor: pointer;
            font-size: 0.85rem;
            transition: background 0.2s;
        }
        .btn-delete:hover { background: #6b3a3a; }
        .empty {
            text-align: center;
            color: #555;
            padding: 40px;
            font-size: 1.1rem;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Task Manager</h1>
        <div class="add-form">
            <input type="text" id="taskInput" placeholder="What needs to be done?" />
            <button onclick="addTask()">Add</button>
        </div>
        <ul class="task-list" id="taskList"></ul>
    </div>
</body>
</html>`
