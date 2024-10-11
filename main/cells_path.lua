cells_pathes = {}

function cells_pathes.add_path(path) 
    if cells_pathes then
        table.insert(cells_pathes, path)
    end
end

function cells_pathes.try_add_cell(cell) 
    local path = cells_pathes.get_last_path();
    if path then 
        for _, value in pairs(path) do
            if value == cell then
                return
            end
        end
        table.insert(path, cell) 
    end
end

function cells_pathes.return_by_cell(cell)
    local path = cells_pathes.get_last_path()
    local cell_index = 1000
    local deleted_cells_index = {}
    for index, value in pairs(path) do
        if value == cell then
            cell_index = index
        end
        if index > cell_index then
            table.insert(deleted_cells_index, index)
        end
    end
    return {path, deleted_cells_index}
end

function cells_pathes.get_last_path()
    return cells_pathes[#cells_pathes]
end

function cells_pathes.clear_path(path) 
    for index, value in pairs(cells_pathes) do
        if value == path then
            cells_pathes[index] = nil
        end
    end
end

return cells_pathes
