import { useGetTodo } from '@/api/generated/todoSchemaDriven';

export default function Test() {
  const { data, error, isPending } = useGetTodo();

  if (isPending) return <div>Pending...</div>;

  if (error) return <div>Error: {error.message}</div>;

  return (
    <div>
      {data?.map((todo) => (
        <div key={todo.title}>{todo.title}:{todo.detail ?? ""}:{todo.progress}</div>
      ))}
    </div>
  );
}
