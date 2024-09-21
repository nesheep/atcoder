defmodule Main do
  def main do
    [r, g, b] = read_ints()
    c = IO.read(:line) |> String.trim()

    ans =
      %{"Red" => r, "Green" => g, "Blue" => b}
      |> Enum.reject(fn {k, _} -> k == c end)
      |> Enum.map(fn {_, v} -> v end)
      |> Enum.min()

    IO.puts(ans)
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
