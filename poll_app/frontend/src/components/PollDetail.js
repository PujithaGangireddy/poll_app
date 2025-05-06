import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";

const dummyPolls = [
  {
    id: 1,
    question: "What is your favorite programming language?",
    options: [
      { id: 11, text: "Go" },
      { id: 12, text: "Python" },
      { id: 13, text: "JavaScript" }
    ]
  },
  {
    id: 2,
    question: "Which OS do you use the most?",
    options: [
      { id: 14, text: "Windows" },
      { id: 15, text: "macOS" },
      { id: 16, text: "Linux" }
    ]
  }
];

const PollDetail = () => {
  const { id } = useParams();
  const [poll, setPoll] = useState(null);

  useEffect(() => {
    const found = dummyPolls.find((p) => p.id === parseInt(id));
    setPoll(found);
  }, [id]);

  if (!poll) return <div className="container mt-4">Poll not found</div>;

  return (
    <div className="container mt-4">
      <h3>{poll.question}</h3>
      <ul className="list-unstyled">
        {poll.options.map((option) => (
          <li key={option.id}>• {option.text}</li>
        ))}
      </ul>
    </div>
  );
};

export default PollDetail;
