defmodule Main do
  def main do
    [h, _w] = read_ints()
    [si, sj] = read_ints()

    c =
      1..h
      |> Enum.map(fn i ->
        IO.read(:line)
        |> String.trim()
        |> String.graphemes()
        |> Enum.with_index(1)
        |> Enum.map(fn {v, j} -> {{i, j}, v == "."} end)
      end)
      |> List.flatten()
      |> Map.new()

    x = IO.read(:line) |> String.trim()

    {ai, aj} =
      x
      |> String.graphemes()
      |> Enum.reduce({si, sj}, fn d, {i, j} = prev ->
        next =
          case d do
            "L" -> {i, j - 1}
            "R" -> {i, j + 1}
            "U" -> {i - 1, j}
            "D" -> {i + 1, j}
          end

        if c[next], do: next, else: prev
      end)

    IO.puts("#{ai} #{aj}")
  end

  def read_ints() do
    IO.read(:line)
    |> String.trim()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end
end
