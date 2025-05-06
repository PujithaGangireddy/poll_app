import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const Polls = () => {
  const [polls, setPolls] = useState([]);

  useEffect(() => {
    // const dummyPolls = [
    //   {
    //     id: 1,
    //     question: "What is your favorite programming language?",
    //     options: [
    //       { id: 11, text: "Go" },
    //       { id: 12, text: "Python" },
    //       { id: 13, text: "JavaScript" }
    //     ]
    //   },
    //   {
    //     id: 2,
    //     question: "Which OS do you use the most?",
    //     options: [
    //       { id: 14, text: "Windows" },
    //       { id: 15, text: "macOS" },
    //       { id: 16, text: "Linux" }
    //     ]
    //   }
    // ];
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    fetch("http://localhost:8080/polls", requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setPolls(data);
      })
      .catch((err) => {
        console.error("Error fetching polls:", err);
      });
  }, []);

  return (
    <div className="container mt-4">
      <div className="text-center">
        <h2>See all the polls you created!</h2>
        <hr />
      </div>
      <table className="table table-bordered table-hover">
        <thead className="thead-dark">
          <tr>
            <th>ID</th>
            <th>Question</th>
          </tr>
        </thead>
        <tbody>
          {polls.map((poll) => (
            <tr key={poll.id}>
              <td>
                <Link to={`/polls/${poll.id}`} className="text-decoration-none">
                  {poll.id}
                </Link>
              </td>
              <td>
                <Link to={`/polls/${poll.id}`} className="text-decoration-none">
                  {poll.question}
                </Link>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Polls;
