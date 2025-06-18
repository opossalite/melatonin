

function ListGroup() {
    let items = [
        'New York',
        'San Francisco',
        'Tokyo',
        'London',
        'Paris',
    ];
    //items = [];
    

    //usually for the key, item should have a .id
    return (
        <>
          <h1>List</h1>
          {items.length === 0 && <p>No item found</p>}
          <ul className="list-group">
            {items.map((item, i) =>
                <li
                  className="list-group-item"
                  key={item}
                  onClick={(event) => console.log(item, i, event)}
                >
                  {item}
                </li>
            )}
          </ul>
        </>
    );
}

export default ListGroup;



