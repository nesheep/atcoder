defmodule Main do
  def main do
    [s, t] = IO.read(:line) |> String.trim() |> String.split(" ")

    ans =
      1..(String.length(s) - 1)
      |> Enum.map(fn w -> {w, 0..(w - 1)} end)
      |> Enum.map(fn {w, cs} ->
        Enum.map(cs, fn c ->
          String.slice(s, c..-1//1)
          |> String.split("", trim: true)
          |> Enum.take_every(w)
          |> Enum.join()
        end)
      end)
      |> List.flatten()
      |> Enum.any?(&(&1 == t))

    IO.puts(if ans, do: "Yes", else: "No")
  end
end
