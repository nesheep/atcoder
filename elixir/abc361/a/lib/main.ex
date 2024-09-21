defmodule Main do
  def main do
    [_, k, x] = read_ints()
    a = read_ints()
    b = a |> List.insert_at(k, x)
    IO.puts(Enum.join(b, " "))
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
