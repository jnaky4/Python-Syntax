from bokeh.plotting import figure, show
from bokeh.models import BoxAnnotation
from bokeh.io import curdoc

# prepare some data
# x_axis = [2, 4, 6, 8, 10]
x_axis = [1, 2, 3, 4, 5]
y_blue = [6, 7, 2, 4, 100]
y_red = [2, 3, 4, 5, 6]
y_dots = [4, 5, 5, 7, 2]

# create a new plot with a title and axis labels
p = figure(title="Multiple glyphs example", x_axis_label="x", y_axis_label="y", sizing_mode="stretch_width", height=1000)

curdoc().theme = "dark_minimal"

# add multiple renderers
p.line(x_axis, y_blue, legend_label="Temp.", line_color="blue", line_width=2)
p.line(x_axis, y_red, legend_label="Rate", line_color="red", line_width=2)
p.circle(x_axis, y_dots, legend_label="Objects", fill_alpha=0.5, fill_color='red', line_color="orange", size=40)
# p.vbar(x_axis, top=y_red, legend_label="Rate", width=0.5, bottom=0, color="red")

# annotations are visual elements that you add to your plot to make it easier to read
low_box = BoxAnnotation(top=20, fill_alpha=0.1, fill_color="red")
mid_box = BoxAnnotation(bottom=20, fill_alpha=0.1, top=80, fill_color="blue")
high_box = BoxAnnotation(bottom=80, fill_alpha=0.1, fill_color="red")

p.add_layout(low_box)
p.add_layout(mid_box)
p.add_layout(high_box)

p.axis.minor_tick_in = -3
p.axis.minor_tick_out = 6

# show the results
show(p)
